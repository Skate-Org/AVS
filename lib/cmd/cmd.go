// Package cmd provides a common utilities and helper function to standarise
// the way omni apps use cobra and viper to produce consistent cli experience
// for both users and devs.
package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd returns a new root cobra command that handles our command line tool.
// It sets up the general viper config and binds the cobra flags to the viper flags.
// It also silences the usage printing when commands error during "running".
func NewRootCmd(appName string, appDescription string, subCmds ...*cobra.Command) *cobra.Command {
	root := &cobra.Command{
		Use:   appName,
		Short: appDescription,
		Args:  cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// NOTE: global prehook for each process
			return nil
		},
	}

	root.AddCommand(subCmds...)

	return root
}

// SilenceErrUsage silences the usage and error printing.
func SilenceErrUsage(cmd *cobra.Command) {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	for _, cmd := range cmd.Commands() {
		SilenceErrUsage(cmd)
	}
}
