package monitor

import (
	"context"
	"time"

	pbCommon "github.com/Skate-Org/AVS/api/pb/common"
	pb "github.com/Skate-Org/AVS/api/pb/relayer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bindingSkateApp "github.com/Skate-Org/AVS/contracts/bindings/SkateApp"
	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/crypto/ecdsa"
	"github.com/Skate-Org/AVS/lib/monitor"
	"github.com/Skate-Org/AVS/lib/on-chain/avs"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"
	"github.com/Skate-Org/AVS/lib/on-chain/network"
	skateappDb "github.com/Skate-Org/AVS/operator/db/skateapp/disk"
)

// WARNING: Must run with wss rpc.
//
// Server subscription only available on websocket endpoints.
func SubscribeSkateApp(addr common.Address, be backend.Backend, ctx context.Context) error {
	contract, err := bindingSkateApp.NewBindingSkateApp(addr, be)
	if err != nil {
		monitor.Logger.Error("Contract binding error: ", "error", err)
		return err
	}

	latest, _ := be.BlockNumber(ctx)
	watchOpts := &bind.WatchOpts{
		Start:   &latest,
		Context: ctx,
	}

	sink := make(chan *bindingSkateApp.BindingSkateAppTaskCreated)
	watcher, err := contract.WatchTaskCreated(watchOpts, sink, nil)
	if err != nil {
		monitor.Logger.Error("Watcher initialization error: ", "error", err)
		return err
	}

	signer := ctx.Value("signer").(*libcmd.SignerConfig)

	var privateKey *ecdsa.PrivateKey
	if signer != nil {
		privateKey, _ = backend.PrivateKeyFromKeystore(common.HexToAddress(signer.Address), signer.Passphrase)
	}

	// Event handler
	go func() {
		for {
			select {
			case task, ok := <-sink:
				if !ok {
					monitor.Logger.Error("Sink error, go-eth related")
					return
				}
				if monitor.Verbose {
					monitor.Logger.Info("Received TaskCreated event:",
						"sender", task.Signer,
						"msg", task.Message,
						"chainId", task.ChainId,
						"chainType", task.ChainType,
						"txHash", task.Raw.TxHash.Hex(),
					)
				}
				if !network.IsSupported(task.ChainType, task.ChainId) {
					monitor.Logger.Info("Unsupported network!", "action", "ignored")
					continue
				}
				PostProcessLog(privateKey, task)
			case err := <-watcher.Err():
				if err != nil && monitor.Verbose {
					monitor.Logger.Error("Watcher received error: ", "error", err)
				}
				return
			}
		}
	}()

	// Wait for the watcher to be closed or an error to occur
	<-watcher.Err()
	return nil
}

func PostProcessLog(privateKey *ecdsa.PrivateKey, bindingTask *bindingSkateApp.BindingSkateAppTaskCreated) error {
	err := dumpLog(bindingTask)
	if err != nil {
		return err
	}
	if privateKey != nil {
		err := signAndBroadcastLog(privateKey, bindingTask)
		if err != nil {
			return err
		}
	}
	return nil
}

func signAndBroadcastLog(privateKey *ecdsa.PrivateKey, bindingTask *bindingSkateApp.BindingSkateAppTaskCreated) error {
	// Step 1: sign the log
	digestHash := avs.TaskDigestHash(
		uint32(bindingTask.TaskId.Int64()), bindingTask.Message, bindingTask.Signer.Hex(),
		pb.ChainType(bindingTask.ChainType), bindingTask.ChainId,
	)
	signature, err := ecdsa.Sign(digestHash, privateKey)
	if err != nil {
		return err
	}

	// Step 2: broad cast log over grpc server
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(timeoutCtx, ":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(),
	)
	defer conn.Close()
	if err != nil {
		if monitor.Verbose {
			monitor.Logger.Fatal("Failed to connect to Relayer", "error", errors.Wrap(err, "signAndBroadcastLog"))
		}
		return err
	}

	client := pb.NewSubmissionClient(conn)

	// Create a new Task
	task := &pb.Task{
		TaskId:    uint32(bindingTask.TaskId.Uint64()),
		Msg:       bindingTask.Message,
		ChainId:   bindingTask.ChainId,
		ChainType: pb.ChainType(bindingTask.ChainType),
		Hash:      bindingTask.TaskHash[:],
		Initiator: bindingTask.Signer.Hex(),
	}

	opAddr := ecdsa.PubkeyToAddress(privateKey.PublicKey).Hex()
	// Create a new SignedMessage
	signedMessage := &pbCommon.OperatorSignature{
		Signature: signature[:],
		Address:   opAddr,
	}

	// Create a new TaskSubmitRequest
	request := &pb.TaskSubmitRequest{
		Task:      task,
		Signature: signedMessage,
	}

	response, err := client.SubmitTask(timeoutCtx, request)
	if err != nil && monitor.Verbose {
		monitor.Logger.Error("Could not submit task", "error", err)
		return err
	}

	if monitor.Verbose {
		monitor.Logger.Info("Response result", "result", response)
	}
	return nil
}

func dumpLog(bindingTask *bindingSkateApp.BindingSkateAppTaskCreated) error {
	err := skateappDb.SkateApp_InsertTask(bindingTask)
	if err != nil && monitor.Verbose {
		monitor.Logger.Info("Can't dump task into db", "error", errors.Wrap(err, "dumpLog"))
	}

	return err
}
