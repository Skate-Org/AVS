package cmd

import (
	"context"

	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/relayer/retrieve"

	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/spf13/cobra"
)

func retrieveCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var configFile string
	var verbose bool
	var enableMetrics bool

	cmd := &cobra.Command{
		Use:   "retrieve",
		Short: "Listen and retrieve task verification signatures from AVS operators",
		Long: `Only accepts operators who registered with Eigenlayer contracts and the Skate AVS.
    Status of the request is not available yet`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger.Info("Relayer: receiving signatures ..")

			config, err := libcmd.ReadConfig[libcmd.EnvironmentConfig]("/environment", configFile)
			if err != nil {
				logger.Fatalf("Can't load config file at %s, error = %v", configFile, err)
				return err
			}
			ctx := context.WithValue(context.Background(), "config", config)

			if enableMetrics {
				logger := logging.NewLoggerWithConsoleWriter()
				metrics := retrieve.NewMetrics(MONITOR_METRICS_PORT, logger)
				ctx = context.WithValue(ctx, "metrics", metrics)
				metrics.Start()
			}

			s := retrieve.NewSubmissionServer(ctx)

			retrieve.Verbose = verbose
			s.Start()

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &configFile)
	libcmd.BindVerbose(cmd, &verbose)
	libcmd.BindMetrics(cmd, &enableMetrics)

	return cmd
}
