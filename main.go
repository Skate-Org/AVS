package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	bindingISkateAVS "github.com/Skate-Org/AVS/contracts/bindings/ISkateAVS"
	"github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"
)

func main() {
	logger := logging.NewLoggerWithConsoleWriter()
	config, _ := cmd.ReadConfig[cmd.EnvironmentConfig]("/environment", "testnet")

	be, _ := backend.NewBackend(config.HttpRPC)
	avs, _ := bindingISkateAVS.NewBindingISkateAVS(common.HexToAddress(config.SkateAVS), be)

	result, _ := avs.Operators(&bind.CallOpts{})

	logger.Info("Skate AVS existing operators", "count", len(result), "addresses", result)
}
