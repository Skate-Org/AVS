package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// NOTE: Redact sensitive log, e.g. authorization header, password, ...
// To be used in the futures
func Redact(input string) string {
	if strings.Contains(input, "token") ||
		strings.Contains(input, "password") ||
		strings.Contains(input, "secret") ||
		strings.Contains(input, "db") ||
		strings.Contains(input, "header") ||
		strings.Contains(input, "key") {
		return "█████"
	}

	return input
}

// verbosity level
func BindVerbose(cmd *cobra.Command, verbose *bool) {
	cmd.Flags().BoolVar(verbose, "verbose", false, "Run with verbose logs")
}

// env config
func BindEnvConfig(cmd *cobra.Command, filename *string) {
	cmd.Flags().StringVar(filename, "config", "testnet", "Config file to set up the environment")
}

func BindSignerConfig(cmd *cobra.Command, filename *string) {
	cmd.Flags().StringVar(filename, "signer-config", "1", "Config file to load the signer")
}

// privateKey as input
// func BindPrivateKey(cmd *cobra.Command, privateKey *string) {
// 	cmd.Flags().StringVar(privateKey, "address", *privateKey, "Run with verbose logs")
// }

// custom signer
func BindSigner(cmd *cobra.Command, signer *string) {
	cmd.Flags().StringVar(signer, "signer-address", *signer, "Override signer address, must be used with an associated passphrase, see `--passphrase` flag")
}

// Passphrase for unlocking signer keystore
func BindPassphrase(cmd *cobra.Command, passphrase *string) {
	cmd.Flags().StringVar(passphrase, "passphrase", *passphrase, "Passphrase to unlock the signer")
}
