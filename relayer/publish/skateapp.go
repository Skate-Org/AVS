package publish

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	pb "skatechain.org/api/pb/relayer"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	bindingSkateGateway "skatechain.org/contracts/bindings/SkateGateway"
	libcmd "skatechain.org/lib/cmd"
	"skatechain.org/lib/crypto/ecdsa"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/avs"
	"skatechain.org/lib/on-chain/backend"
	"skatechain.org/relayer/db/skateapp/disk"
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

	for key, taskGroup := range taskGroups {
		if len(taskGroup)*10_000 >= operatorCount*6_666 {
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
		}
	}

	if len(batchTaskId) == 0 {
		return
	}

	chainId := new(big.Int).SetUint64(config.MainChainId)
	avsTransactor, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

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

	for key, taskGroup := range taskGroups {
		completedTask := disk.CompletedTask{
			TaskId:    key.TaskId,
			ChainId:   key.ChainId,
			ChainType: key.ChainType,
		}
		// TODO: publish tasks logic to all supported Skate Gateway,
		// improve environement configured logic
		switch key.ChainId {
		case 421614:
			task := taskGroup[0]
			be, _ := backend.NewBackend("https://arbitrum-sepolia.blockpi.network/v1/rpc/public")
			gatewayAddress := common.HexToAddress("0xc1Eb0ffdb88c59A043ab5B4fBf200795Cd5Acd03")
			gatewayContract, _ := bindingSkateGateway.NewBindingSkateGateway(gatewayAddress, be)
			transactor, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(421614))

			relayerLogger.Info("Submitting message to Gateway", "chainID", key.ChainId)
			tx, _ := gatewayContract.PostMsg(
				transactor,
				big.NewInt(int64(key.TaskId)),
				task.Message,
				common.HexToAddress(task.Initiator),
			)
			receipt, _ := backend.WaitMined(context.Background(), &be, tx)
			relayerLogger.Info("Submitted to gateway, receipt:", "status", receipt.Status, "chainID", key.ChainId)
		}
		disk.InsertCompletedTask(completedTask)
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
