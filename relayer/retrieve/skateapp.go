package retrieve

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"github.com/Skate-Org/AVS/api"
	pb "github.com/Skate-Org/AVS/api/pb/relayer"
	bindingISkateAVS "github.com/Skate-Org/AVS/contracts/bindings/ISkateAVS"
	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/crypto/ecdsa"
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/on-chain/avs"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"
	avsMemcache "github.com/Skate-Org/AVS/relayer/db/avs/mem"
	skateappDb "github.com/Skate-Org/AVS/relayer/db/skateapp/disk"
	skateappMemcache "github.com/Skate-Org/AVS/relayer/db/skateapp/mem"
)

var (
	retrieveLogger = logging.NewLoggerWithConsoleWriter()
	Verbose        = true
	taskCache      = skateappMemcache.NewCache(100 * 1024 * 1024) // 100MB
	operatorCache  = avsMemcache.NewCache(2 * 1024 * 1024)        // 2MB
)

type submissionServer struct {
	pb.UnimplementedSubmissionServer
	ctx context.Context
}

func NewSubmissionServer(ctx context.Context) *submissionServer {
	return &submissionServer{
		ctx: ctx,
	}
}

func (s *submissionServer) Start() {
	grpc_server := grpc.NewServer()

	pb.RegisterSubmissionServer(grpc_server, s)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	retrieveLogger.Info("Server listening", "Address", lis.Addr().String(), "network", lis.Addr().Network())
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *submissionServer) SubmitTask(_ context.Context, in *pb.TaskSubmitRequest) (*pb.TaskSubmitReply, error) {
	config := s.ctx.Value("config").(*libcmd.EnvironmentConfig)
	if Verbose {
		retrieveLogger.Info("Got request", "payload", in)
	}

	// Step 1: Verify the operator
	be, _ := backend.NewBackend(config.HttpRPC)
	avsContract, _ := bindingISkateAVS.NewBindingISkateAVS(
		common.HexToAddress(config.SkateAVS), be,
	)
	isValidOperator, err := isOperator(avsContract, in.Signature.Address)
	if err != nil {
		if Verbose {
			retrieveLogger.Error("Validator address format error", "error", err)
		}
		return &pb.TaskSubmitReply{
			Result: pb.TaskStatus_REJECTED,
		}, api.NewInvalidArgError("signer address format error")
	}
	if !isValidOperator {
		if Verbose {
			retrieveLogger.Error("Not an operator", "address", in.Signature.Address)
		}
		return &pb.TaskSubmitReply{
			Result: pb.TaskStatus_REJECTED,
		}, api.NewInvalidArgError(fmt.Sprintf("%s is not a Skate AVS operator", in.Signature.Address))
	}

	// Step 2: Verify signature
	signature := [65]byte(in.Signature.Signature)
	taskDigest := avs.TaskDigestHash(in.Task.TaskId, in.Task.Msg, in.Task.Initiator, in.Task.ChainType, in.Task.ChainId)
	valid, err := ecdsa.Verify(
		common.HexToAddress(in.Signature.Address),
		taskDigest,
		signature,
	)
	if err != nil {
		if Verbose {
			retrieveLogger.Error("Signature format error", "error", err)
		}
		return &pb.TaskSubmitReply{
			Result: pb.TaskStatus_REJECTED,
		}, api.NewInvalidArgError("Signature format error, must be 65 bytes")
	}
	if !valid {
		if Verbose {
			retrieveLogger.Error("Signature is invalid",
				"operator", in.Signature.Address,
				"signature", signature,
				"TaskId", in.Task.TaskId,
				"ChainType", in.Task.ChainType,
				"ChainId", in.Task.ChainId,
			)
		}
		return &pb.TaskSubmitReply{
			Result: pb.TaskStatus_REJECTED,
		}, api.NewInvalidArgError("Invalid signature")
	}

	// Step 3: Update the db and push to memcache
	msg := skateappMemcache.Message{
		TaskId:    in.Task.TaskId,
		ChainId:   in.Task.ChainId,
		ChainType: in.Task.ChainType,
		Message:   in.Task.Msg,
		Initiator: in.Task.Initiator,
	}

	msgKey := skateappMemcache.GenKey(msg)
	taskCache.CacheMessage(msgKey, msg)
	sig := skateappMemcache.Signature{
		Operator:  in.Signature.Address,
		Signature: signature,
	}
	taskCache.AppendSignature(msgKey, sig)

	signedTask := skateappDb.SignedTask{
		TaskId:    in.Task.TaskId,
		Message:   in.Task.Msg,
		Initiator: in.Task.Initiator,
		ChainId:   in.Task.ChainId,
		ChainType: uint32(in.Task.ChainType),
		Hash:      in.Task.Hash,
		Operator:  in.Signature.Address,
		Signature: in.Signature.Signature,
	}
	err = skateappDb.InsertSignedTask(signedTask)
	if err != nil && Verbose {
		retrieveLogger.Error("Insert signed task to db failed", "error", err)
		return &pb.TaskSubmitReply{
			Result: pb.TaskStatus_REJECTED,
		}, api.NewInternalError("Server can't securely store signed task object")
	}

	return &pb.TaskSubmitReply{
		Result: pb.TaskStatus_PROCESSING,
	}, nil
}

// NOTE: Right now we control the operators list,
// therefore cache revalidation period is set to INF (no expiration).
// Might need to change in the future.
func isOperator(avsContract *bindingISkateAVS.BindingISkateAVS, address string) (bool, error) {
	// step 1: look up cache
	cachedOperator, _ := operatorCache.GetOperator(address)
	if cachedOperator != nil {
		return true, nil
	}

	// step 2: populate cache with on-chain data
	operators, err := avsContract.Operators(&bind.CallOpts{})
	if err != nil {
		return false, errors.Wrap(err, "isOperator.Operators")
	}

	operatorCache.CacheOperatorCount(uint32(len(operators)))
	for _, op := range operators {
		cacheOp := avsMemcache.Operator{
			Address: op.Addr.Hex(),
		}
		operatorCache.CacheOperator(cacheOp)
		if op.Addr.Hex() == address {
			return true, nil
		}
		// TODO: cache stake amounts as well
	}

	return false, nil
}
