package backend

import (
	elEthClient "github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	elTxMgr "github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"skatechain.org/lib/logging"
)

// NOTE: add Skate specific logic for multichain ops in future versions,
type Backend struct {
	elEthClient.Client
	RPC string
}

func NewBackend(rpc string) (Backend, error) {
	elClient, err := elEthClient.NewClient(rpc)
	return Backend{Client: elClient, RPC: rpc}, err
}

var _ elEthClient.Client = (*Backend)(nil)

type TxManager struct {
	elTxMgr.TxManager
}

var _ elTxMgr.TxManager = (*TxManager)(nil)

func NewSimpleTxManager(wallet wallet.Wallet, backend Backend, sender common.Address) *TxManager {
	logger := logging.NewLoggerWithConsoleWriter()
	return &TxManager{
		TxManager: elTxMgr.NewSimpleTxManager(wallet, backend, logger, sender),
	}
}
