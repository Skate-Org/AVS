package monitor

import (
	"context"
	// "crypto/tls"
	"time"

	pbCommon "github.com/Skate-Org/AVS/api/pb/common"
	pb "github.com/Skate-Org/AVS/api/pb/relayer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
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

	var metrics *Metrics
	if ctx.Value("metrics") != nil {
		metrics = ctx.Value("metrics").(*Metrics)
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

				IncreaseTaskProcessed(metrics, TaskStatus_DETECTED)
				PostProcessLog(privateKey, task, metrics)
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

func PostProcessLog(privateKey *ecdsa.PrivateKey, bindingTask *bindingSkateApp.BindingSkateAppTaskCreated, metrics *Metrics) error {
	err := dumpLog(bindingTask)
	if err != nil {
		IncreaseTaskProcessed(metrics, TaskStatus_SAVE_FAILED)
		return err
	}
	IncreaseTaskProcessed(metrics, TaskStatus_SAVED)
	if privateKey != nil {
		err := signAndBroadcastLog(privateKey, bindingTask)
		if err != nil {
			IncreaseTaskProcessed(metrics, TaskStatus_VERIFY_FAILED)
			return err
		}
		IncreaseTaskProcessed(metrics, TaskStatus_VERIFIED)
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

	// tlsConfig := &tls.Config{
	// 	ServerName: "relayer.skatechain.org",
	// }
	// creds := credentials.NewTLS(tlsConfig)
	// if monitor.Verbose {
	// 	monitor.Logger.Info("Dialing relayer.skatechain.org ...")
	// }
	// conn, err := grpc.DialContext(timeoutCtx, "relayer.skatechain.org:443",
	// 	grpc.WithTransportCredentials(creds),
	// 	grpc.WithBlock(),
	// )

	conn, err := grpc.DialContext(timeoutCtx, ":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	monitor.Logger.Info("Sending to relayer...")
	if err != nil {
		monitor.Logger.Fatal("Failed to connect to Relayer at relayer.skatechain.org", "error", errors.Wrap(err, "signAndBroadcastLog"))
		return err
	}
	if monitor.Verbose {
		monitor.Logger.Info("Connected!", "conn", conn.GetState().String())
	}
	defer conn.Close()

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

	if monitor.Verbose {
		monitor.Logger.Info("Submitting signed task ... ", "payload", request)
	}
	response, err := client.SubmitTask(timeoutCtx, request)
	if err != nil && monitor.Verbose {
		monitor.Logger.Error("Submission failed: ", "error", err)
		return err
	}

	if monitor.Verbose {
		monitor.Logger.Info("Submission approved: ", "result", response.String())
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
