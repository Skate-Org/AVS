package backend

import (
	"context"
	"time"

	elEthClient "github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
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

var logger = logging.NewLoggerWithConsoleWriter()

const MAX_WAIT = 10

// WaitMined waits for tx to be mined on the blockchain.
// Returns the transaction receipt when the transaction is mined.
func (be *Backend) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second * 6)
	defer queryTicker.Stop()

	count := 0

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
			receipt, err := be.TransactionReceipt(ctx, tx.Hash())
			logger.Info("Error", "error", err, "hash", tx.Hash())
			if receipt != nil {
				return receipt, nil
			}
		}
		count += 1
		if count >= MAX_WAIT {
			return nil, errors.New("Wait timed out!")
		}
	}
}
