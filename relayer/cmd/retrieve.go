package cmd

import (
	"context"

	"skatechain.org/lib/logging"
	"skatechain.org/relayer/retrieve"

	"github.com/spf13/cobra"
	libcmd "skatechain.org/lib/cmd"
)

func retrieveCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var configFile string

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

			s.Start()

			return nil
		},
	}

	libcmd.BindEnvConfig(cmd, &configFile)

	verbose := true
	libcmd.BindVerbose(cmd, &verbose)
	if !verbose {
		retrieve.Verbose = false
	}

	return cmd
}
