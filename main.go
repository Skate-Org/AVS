package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	bindingISkateAVS "skatechain.org/contracts/bindings/ISkateAVS"
	"skatechain.org/lib/cmd"
	"skatechain.org/lib/logging"
	"skatechain.org/lib/on-chain/backend"
)

func main() {
	logger := logging.NewLoggerWithConsoleWriter()
	config, _ := cmd.ReadConfig[cmd.EnvironmentConfig]("/environment", "testnet")

	be, _ := backend.NewBackend(config.HttpRPC)
	avs, _ := bindingISkateAVS.NewBindingISkateAVS(common.HexToAddress(config.SkateAVS), be)

	result, _ := avs.Operators(&bind.CallOpts{})

	logger.Info("Skate AVS existing operators", "count", len(result), "addresses", result)
}
