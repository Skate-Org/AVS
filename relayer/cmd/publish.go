package cmd

import (
	"context"
	"fmt"

	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"
	"github.com/Skate-Org/AVS/relayer/publish"

	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

const MONITOR_METRICS_PORT = "9001"

func publishCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var envConfigFile string
	var signerConfigFile string
	var overrideSigner string
	var passphrase string
	var verbose bool
	var enableMetrics bool

	cmd := &cobra.Command{
		Use:   "publish",
		Short: "Publish verified quorums and relay messages",
		Long:  `Publish verified quorums to skate AVS, then relay messages in respective gateway contracts. Relayer slashing is not yet implemented`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			envConfig, err := libcmd.ReadConfig[libcmd.EnvironmentConfig]("/environment", envConfigFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", envConfigFile, err)
				return err
			}
			ctx := context.WithValue(context.Background(), "config", envConfig)

			signerConfig, err := libcmd.ReadConfig[libcmd.SignerConfig]("/signer/relayer", signerConfigFile)
			if overrideSigner != "" {
				signerConfig.Address = overrideSigner
			}
			if passphrase != "" {
				signerConfig.Passphrase = passphrase
			}

			if signerConfig.Address == "" {
				logger.Fatal("No signer provided")
			}
			_, err = backend.PrivateKeyFromKeystore(common.HexToAddress(signerConfig.Address), signerConfig.Passphrase)
			if err != nil {
				logger.Fatal("Invalid keystore for relayer", signerConfig)
				return err
			}

			if enableMetrics {
				logger := logging.NewLoggerWithConsoleWriter()
				metrics := publish.NewMetrics(MONITOR_METRICS_PORT, logger)
				ctx = context.WithValue(ctx, "metrics", metrics)
				metrics.Start()
			}

			logger.Info("Relayer: Ready to publish tasks to AVS ..",
				"signer", signerConfig.Address,
				"fromConfig", fmt.Sprintf("configs/relayer/%s.yaml", signerConfigFile),
			)
			ctx = context.WithValue(ctx, "signer", signerConfig)

			publish.Verbose = verbose
			publish.PublishTaskToAVSAndGateway(ctx)

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &envConfigFile)
	libcmd.BindSignerConfig(cmd, &signerConfigFile)
	libcmd.BindSigner(cmd, &overrideSigner)
	libcmd.BindPassphrase(cmd, &passphrase)
	libcmd.BindVerbose(cmd, &verbose)
	libcmd.BindMetrics(cmd, &enableMetrics)

	return cmd
}
