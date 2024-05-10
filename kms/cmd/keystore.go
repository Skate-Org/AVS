package cmd

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/on-chain/backend"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func storePrivatekeyCmd() *cobra.Command {
	logger := logging.NewLoggerWithConsoleWriter()

	var privateKey string
	var passphrase string
	var savePath string

	cmd := &cobra.Command{
		Use:   "store",
		Short: "Store encrypted private key with password guarded",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if privateKey == "" {
				logger.Error("private key not provided")
				return errors.New("Private key not provided")
			}
			if passphrase == "" {
				logger.Error("Pass phrase not provided")
				return errors.New("Pass phrase not provided")
			}
			if len(passphrase) < 8 {
				logger.Error("Pass phrase should have at least 8 characters")
				return errors.New("Pass phrase should have at least 8 characters")
			}
			address := backend.DumpECDSAPrivateKeyToKeystore(privateKey, passphrase)
			logger.Info("Private key successfully encrypted in `keystore`")

			if savePath != "" {
				// Building the final path with the provided savePath
				finalPath := filepath.Join("configs", "signer", savePath+".yaml")
				config := map[string]string{
					"address":    address.Hex(),
					"passphrase": passphrase,
				}
				data, err := yaml.Marshal(config)
				if err != nil {
					logger.Error("Failed to marshal YAML", "error", err)
					return errors.Wrap(err, "Failed to marshal YAML")
				}
				if err := os.WriteFile(finalPath, data, 0o644); err != nil {
					logger.Error("Failed to write YAML file", "error", err)
					return errors.Wrap(err, "Failed to write YAML file")
				}
				logger.Info("Configuration saved", "path", finalPath)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&privateKey, "private-key", "k", "", "Private key")
	cmd.Flags().StringVarP(&passphrase, "passphrase", "p", "", "Passphrase")
	cmd.Flags().StringVarP(&savePath, "save-path", "s", "", "File to save the signer config after create")

	return cmd
}
