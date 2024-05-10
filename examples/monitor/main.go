package main

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Skate-Org/AVS/contracts/bindings/IERC20"
	"github.com/Skate-Org/AVS/lib/monitor"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"
	"github.com/Skate-Org/AVS/lib/on-chain/network"
)

// NOTE: example of USDC monitor service on [mainnet, polygon]
func main() {
	mainnet := network.ChainID(1)
	mainnet_backend0, _ := backend.NewBackend("wss://ethereum-rpc.publicnode.com")
	mainnet_USDC := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	// polygon := network.ChainID(137)
	// polygon_backend0, _ := backend.NewBackend("wss://polygon-bor-rpc.publicnode.com")
	// polygon_USDC := common.HexToAddress("0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359")

	contractAddrs := map[network.ChainID]common.Address{
		mainnet: mainnet_USDC,
		// polygon: polygon_USDC,
	}

	backends := map[network.ChainID][]backend.Backend{
		mainnet: {mainnet_backend0},
		// polygon: {polygon_backend0},
	}

	ctx := map[network.ChainID]context.Context{
		mainnet: context.Background(),
		// polygon: context.Background(),
	}

	println("Example retrieving USDC transfer on mainnet...")
	monitor := monitor.NewMonitor(ctx, contractAddrs, backends)
	monitor.Start(pollLog)
}

// NOTE: example subscribe monitor
func subscribeLog(addr common.Address, backend backend.Backend, ctx context.Context) error {
	contract, err := bindingIERC20.NewBindingIERC20(addr, backend)
	if err != nil {
		monitor.Logger.Error("Contract binding error: ", "error", err)
		return err
	}

	latest, _ := backend.BlockNumber(ctx)
	watchOpts := &bind.WatchOpts{
		Start:   &latest,
		Context: ctx,
	}

	sink := make(chan *bindingIERC20.BindingIERC20Transfer)
	watcher, err := contract.WatchTransfer(watchOpts, sink, nil, nil)
	if err != nil {
		if monitor.Verbose {
			monitor.Logger.Error("Watcher initialization error: ", "error", err)
		}
		return err
	}

	// Event handler
	go func() {
		for {
			select {
			case event, ok := <-sink:
				if !ok {
					return
				}
				if monitor.Verbose {
					monitor.Logger.Info("Received Transfer event:", "from", event.From, "to", event.To, "value", event.Value, "rpc", backend.RPC)
				}
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

// NOTE: example polling monitor
func pollLog(addr common.Address, backend backend.Backend, ctx context.Context) error {
	contract, err := bindingIERC20.NewBindingIERC20(addr, backend)
	if err != nil {
		monitor.Logger.Error("Contract binding error: ", "error", err)
		return err
	}

	// NOTE: Polling interval = 12s
	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()

	// Retrieve the latest block number as the starting point for the filter
	latest, err := backend.BlockNumber(ctx)
	if err != nil {
		if monitor.Verbose {
			monitor.Logger.Error("Error retrieving latest block number: ", "error", err)
		}
		return err
	}

	// Filter options
	queryOpts := &bind.FilterOpts{
		Start:   latest,
		Context: ctx,
	}

	// Function to process events
	processEvents := func() error {
		it, err := contract.FilterTransfer(queryOpts, nil, nil)
		if err != nil {
			if monitor.Verbose {
				monitor.Logger.Error("Error fetching events: ", "error", err)
			}
			return err
		}

		// Process all events since the last poll
		for it.Next() {
			event := it.Event
			if monitor.Verbose {
				monitor.Logger.Info("Polled Transfer event:", "from", event.From, "to", event.To, "amount", event.Value, "rpc", backend.RPC)
			}
		}

		// Update the starting block for the next query to be the block number of the last fetched event
		if it.Event != nil {
			queryOpts.Start = it.Event.Raw.BlockNumber
		}
		return nil
	}

	// Immediately process events once before starting the ticker
	if err := processEvents(); err != nil {
		return err
	}

	// Polling loop
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := processEvents(); err != nil {
					continue
				}
			case <-ctx.Done():
				if monitor.Verbose {
					monitor.Logger.Info("Polling stopped due to context cancellation")
				}
				return
			}
		}
	}()
	// Keep the function alive until the context is cancelled
	<-ctx.Done()
	return nil
}
