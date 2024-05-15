package cmd

import (
	"context"

	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/relayer/retrieve"

	"github.com/spf13/cobra"
	libcmd "github.com/Skate-Org/AVS/lib/cmd"
)

func retrieveCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var configFile string
  var verbose bool

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
			s := retrieve.NewSubmissionServer(ctx)

		retrieve.Verbose = verbose
			s.Start()

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &configFile)
	libcmd.BindVerbose(cmd, &verbose)

	return cmd
}
