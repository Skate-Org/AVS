package backend

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/Skate-Org/AVS/lib/logging"
)

var logger = logging.NewLoggerWithConsoleWriter()

const MAX_WAIT = 10

// WaitMined waits for tx to be mined on the blockchain.
// Returns the transaction receipt when the transaction is mined.
func WaitMined(ctx context.Context, backend *Backend, tx *types.Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second * 6)
	defer queryTicker.Stop()

	count := 0

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
			receipt, err := backend.TransactionReceipt(ctx, tx.Hash())
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
