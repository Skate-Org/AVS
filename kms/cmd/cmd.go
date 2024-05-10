package cmd

import (
	"github.com/spf13/cobra"
	"github.com/Skate-Org/AVS/lib/buildinfo"
	libcmd "github.com/Skate-Org/AVS/lib/cmd"
)

const VERSION = "v1.0.0"

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"Key management system",
		"CLI for Private Key management when intereacting with Skate App",
		storePrivatekeyCmd(),
		buildinfo.BuildInfoCmd(VERSION), // TODO: seperate package info
	)
}
