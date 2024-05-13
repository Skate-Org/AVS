package publish

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	pb "github.com/Skate-Org/AVS/api/pb/relayer"
	bindingISkateAVS "github.com/Skate-Org/AVS/contracts/bindings/ISkateAVS"
	bindingSkateGateway "github.com/Skate-Org/AVS/contracts/bindings/SkateGateway"
	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/crypto/ecdsa"
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/on-chain/avs"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"
	"github.com/Skate-Org/AVS/relayer/db/skateapp/disk"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var relayerLogger = logging.NewLoggerWithConsoleWriter()

func PublishTaskToAVSAndGateway(ctx context.Context) {
	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()

	config := ctx.Value("config").(*libcmd.EnvironmentConfig)
	be, err := backend.NewBackend(config.HttpRPC)
	if err != nil {
		relayerLogger.Fatal("AVS rpc error", "rpcUrl", config.HttpRPC)
		return
	}
	avsContract, err := bindingISkateAVS.NewBindingISkateAVS(
		common.HexToAddress(config.SkateAVS), be,
	)
	if err != nil {
		relayerLogger.Fatal("Invalid avs contract", "address", config.SkateAVS)
		return
	}

	signer := ctx.Value("signer").(*libcmd.SignerConfig)
	privateKey, _ := backend.PrivateKeyFromKeystore(common.HexToAddress(signer.Address), signer.Passphrase)

	// Call submitTasks immediately
	submitTasksToAvs(avsContract, &be, config, privateKey)

	for {
		select {
		case <-ctx.Done():
			relayerLogger.Info("AVS publish process stopped!")
			return
		case <-ticker.C:
			submitTasksToAvs(avsContract, &be, config, privateKey)
		}
	}
}

// TaskGroupKey is a struct to hold the key for grouping tasks
type TaskGroupKey struct {
	TaskId    uint32
	ChainId   uint32
	ChainType uint32
}

// TaskGroup is a struct to hold a group of tasks
type (
	TaskGroup      = []disk.SignedTask
	SignatureTuple = bindingISkateAVS.ISkateAVSSignatureTuple
)

type VerifiedTask struct {
	TaskId    uint32
	ChainId   uint32
	ChainType uint32
	Initiator string
	Message   string
}

func submitTasksToAvs(avsContract *bindingISkateAVS.BindingISkateAVS, be *backend.Backend, config *libcmd.EnvironmentConfig, privateKey *ecdsa.PrivateKey) {
	batchTaskId := make([]*big.Int, 0)
	batchMessageData := make([][]byte, 0)
	batchSignatureTuples := make([][]SignatureTuple, 0)
	operators, _ := avsContract.Operators(&bind.CallOpts{})
	operatorCount := len(operators)

	// Fetch pending tasks
	tasks, err := fetchPendingTasks()
	if err != nil {
		relayerLogger.Error("Failed to fetch pending tasks", "error", err)
		return
	}

	// Group tasks by (taskId, chainId, chainType)
	taskGroups := make(map[TaskGroupKey]TaskGroup)
	for _, task := range tasks {
		key := TaskGroupKey{TaskId: task.TaskId, ChainId: task.ChainId, ChainType: task.ChainType}
		if group, exists := taskGroups[key]; exists {
			group = append(group, task)
			taskGroups[key] = group
		} else {
			taskGroups[key] = []disk.SignedTask{task}
		}
	}

	verifiedTasks := make([]VerifiedTask, 0)

	// Step 1: Fitler those tasks with quorum threshold reached
	for key, taskGroup := range taskGroups {
		// NOTE: Check for BFT consensus reached.
		quorumReached := len(taskGroup)*10_000 >= operatorCount*6_666
		if !quorumReached {
			continue
		}

		if Verbose {
			relayerLogger.Info("Task approved for submission", "task key", key)
		}
		taskId := new(big.Int).SetUint64(uint64(key.TaskId))
		task := taskGroup[0]
		messageData := avs.TaskData(task.Message, task.Initiator, pb.ChainType_EVM, task.ChainId)
		batchTaskId = append(batchTaskId, taskId)
		batchMessageData = append(batchMessageData, messageData)

		signatureTuples := make([]SignatureTuple, len(taskGroup))
		for index, task := range taskGroup {
			signatureTuples[index] = SignatureTuple{
				Operator:  common.HexToAddress(task.Operator),
				Signature: task.Signature,
			}
		}
		// signatures must be sorted by address
		sort.Slice(signatureTuples, func(i, j int) bool {
			return signatureTuples[i].Operator.Big().Cmp(signatureTuples[j].Operator.Big()) < 0
		})

		// NOTE: hack to ensure tuple is unique, should be filter at db level
		// Prevent the case where multiple process with the same operator running.
		var uniqueSignatureTuples []SignatureTuple
		uniqueSignatureTuples = append(uniqueSignatureTuples, signatureTuples[0]) // 100% len > 0

		for i := 1; i < len(signatureTuples); i++ {
			if signatureTuples[i].Operator.Big().Cmp(uniqueSignatureTuples[len(uniqueSignatureTuples)-1].Operator.Big()) != 0 {
				uniqueSignatureTuples = append(uniqueSignatureTuples, signatureTuples[i])
			}
		}
		batchSignatureTuples = append(batchSignatureTuples, uniqueSignatureTuples)

		// NOTE: insert this verified task for subsequent publishing process.
		verifiedTasks = append(verifiedTasks, VerifiedTask{
			TaskId:    task.TaskId,
			ChainId:   task.ChainId,
			ChainType: task.ChainType,
			Initiator: task.Initiator,
			Message:   task.Message,
		})
	}

	if len(batchTaskId) == 0 {
		return
	}

	// Step 2: Publish batch verified tasks to the AVS
	chainId := new(big.Int).SetUint64(config.MainChainId)
	avsTransactor, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	// Step 2.1: simulate
	transactorNoSend := *avsTransactor
	transactorNoSend.NoSend = true
	_, err = avsContract.BatchSubmitData(
		&transactorNoSend,
		batchTaskId,
		batchMessageData,
		batchSignatureTuples,
	)
	if err != nil {
		relayerLogger.Error("Transaction simulation failed", "error", errors.Wrap(err, "SkateAVS.BatchSubmitData"))
		return
	}

	// Step 2.2: call batchSubmitData on AVS
	if Verbose {
		relayerLogger.Info("Submitting batched tasks to Skate AVS ..")
	}
	tx, err := avsContract.BatchSubmitData(
		avsTransactor,
		batchTaskId,
		batchMessageData,
		batchSignatureTuples,
	)
	if err != nil {
		relayerLogger.Error("Failed to submit transaction", "error", errors.Wrap(err, "SkateAVS.BatchSubmitData"))
		return
	}
	if Verbose {
		relayerLogger.Info("Verification request sent", "txHash", tx.Hash().Hex())
	}
	receipt, err := backend.WaitMined(context.Background(), be, tx)
	if err != nil {
		relayerLogger.Error("Failed to get transaction receipt", "error", err)
		return
	}
	if Verbose {
		relayerLogger.Info("Transaction receipt: ", "status", receipt.Status, "gasUsed", receipt.GasUsed, "gasPrice", receipt.EffectiveGasPrice.Uint64())
	}

	// Step 3: publish verified tasks to the gateway
	TASK_PUBLISHED := false
	for _, verifiedTask := range verifiedTasks {
		// Step 3.1: Publish to destination chain
		completedTask := disk.CompletedTask{
			TaskId:    verifiedTask.TaskId,
			ChainId:   verifiedTask.ChainId,
			ChainType: verifiedTask.ChainType,
		}

		switch pb.ChainType(verifiedTask.ChainType) {
		case pb.ChainType_EVM:
			switch verifiedTask.ChainId {
			case 421614:
				be, err := backend.NewBackend("https://arbitrum-sepolia.blockpi.network/v1/rpc/public")
				if err != nil {
					continue
				}
				gatewayAddress := common.HexToAddress("0xc1Eb0ffdb88c59A043ab5B4fBf200795Cd5Acd03")
				gatewayContract, _ := bindingSkateGateway.NewBindingSkateGateway(gatewayAddress, be)
				transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(421614))
				if err != nil {
					continue
				}

				relayerLogger.Info("Submitting message to Gateway", "chainID", verifiedTask.ChainId)
				tx, err := gatewayContract.PostMsg(
					transactor,
					big.NewInt(int64(verifiedTask.TaskId)),
					verifiedTask.Message,
					common.HexToAddress(verifiedTask.Initiator),
				)
				if err != nil {
					continue
				}
				receipt, err := backend.WaitMined(context.Background(), &be, tx) // NOTE: can be rpc error, regardless need to re-publish next time
				if err != nil {
					continue
				}
				relayerLogger.Info("Submitted to gateway, receipt:", "status", receipt.Status, "chainID", verifiedTask.ChainId)

				TASK_PUBLISHED = true
			}
			// TODO: integrate more chains
		case pb.ChainType_SOLANA:
			// TODO:

		default:
			relayerLogger.Error("Unsupported chain type, ignored")
		}

		// Step 3.2: Cache completed entry in the db
		if TASK_PUBLISHED {
			disk.InsertCompletedTask(completedTask)
		}
	}
}

func fetchPendingTasks() ([]disk.SignedTask, error) {
	query := fmt.Sprintf(`
    SELECT *
    FROM %s s
    WHERE NOT EXISTS (
        SELECT 1 FROM %s c
        WHERE c.taskId = s.taskId AND c.chainId = s.chainId AND c.chainType = s.chainType
    )
  `, disk.SignedTaskSchema, disk.CompletedTaskSchema)
	rows, err := disk.SkateAppDB.Query(query)

	var pendingTasks []disk.SignedTask
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task disk.SignedTask
		var entryid int

		err := rows.Scan(
			&entryid, &task.TaskId, &task.Message, &task.Initiator,
			&task.ChainId, &task.ChainType, &task.Hash, &task.Operator, &task.Signature,
		)
		if err != nil {
			return nil, err
		}
		pendingTasks = append(pendingTasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return pendingTasks, nil
}
