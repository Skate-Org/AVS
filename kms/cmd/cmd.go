package cmd

import (
	"github.com/spf13/cobra"
	"skatechain.org/lib/buildinfo"
	libcmd "skatechain.org/lib/cmd"
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
