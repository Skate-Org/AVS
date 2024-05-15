package cmd

import (
	"context"
	"fmt"

	"github.com/Skate-Org/AVS/lib/logging"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/monitor"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"
	"github.com/Skate-Org/AVS/lib/on-chain/network"
	skateappDb "github.com/Skate-Org/AVS/operator/db/skateapp/disk"
	operatorMonitor "github.com/Skate-Org/AVS/operator/monitor"
)

// TODO: decouple monitor from signing service
func monitorSkateAppCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var envConfigFile string
	var signerConfigFile string
	var overrideSigner string
	var passphrase string
	var verbose bool

	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor TaskCreated events from Skate AVS, verify and dispatch to relayer",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			envConfig, err := libcmd.ReadConfig[libcmd.EnvironmentConfig]("/environment", envConfigFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", envConfigFile, err)
				return err
			}
			ctx := context.WithValue(context.Background(), "config", envConfig)

			signerConfig, err := libcmd.ReadConfig[libcmd.SignerConfig]("/signer/operator", signerConfigFile)
			if overrideSigner != "" {
				signerConfig.Address = overrideSigner
			}
			if passphrase != "" {
				signerConfig.Passphrase = passphrase
			}

			if signerConfig.Address != "" {
				ctx = context.WithValue(ctx, "signer", signerConfig)

				_, err := backend.PrivateKeyFromKeystore(common.HexToAddress(signerConfig.Address), signerConfig.Passphrase)
				if err != nil {
					logger.Fatal("Invalid keystore for signer", "configFile", signerConfig)
					return err
				}
				logger.Info("Operator: monitoring and processing tasks ..",
					"signer", signerConfig.Address,
					"fromConfig", fmt.Sprintf("configs/signer/operator/%s.yaml", signerConfigFile),
				)
			} else {
				logger.Info("No signer provided, running with watch-only mode...")
			}

			monitor.Verbose = verbose
			startMonitor(ctx)

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &envConfigFile)
	libcmd.BindSignerConfig(cmd, &signerConfigFile)
	libcmd.BindSigner(cmd, &overrideSigner)
	libcmd.BindPassphrase(cmd, &passphrase)
	libcmd.BindVerbose(cmd, &verbose)

	return cmd
}

// TODO: populate context to the runner
func startMonitor(ctx context.Context) {
	env := ctx.Value("config").(*libcmd.EnvironmentConfig)
	nollie := network.ChainID(env.SkateChainId)
	nollie_backend0, _ := backend.NewBackend(env.SkateWSSRPC)
	nollie_SkateApp := common.HexToAddress(env.SkateApp)

	skateappDb.InitializeSkateApp()

	contractAddrs := map[network.ChainID]common.Address{
		nollie: nollie_SkateApp,
	}

	backends := map[network.ChainID][]backend.Backend{
		nollie: {nollie_backend0},
	}

	ctxs := map[network.ChainID]context.Context{
		nollie: ctx,
	}

	monitor := monitor.NewMonitor(ctxs, contractAddrs, backends)
	monitor.Start(operatorMonitor.SubscribeSkateApp)
}
