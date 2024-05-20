package cmd

import (
	"github.com/spf13/cobra"
	"github.com/Skate-Org/AVS/lib/buildinfo"
	libcmd "github.com/Skate-Org/AVS/lib/cmd"
)

const VERSION = "1.1.0"

// New returns a new root cobra command that handles our command line tool.
func New() *cobra.Command {
	return libcmd.NewRootCmd(
		"Key management service",
		"CLI for Key management service when intereacting with Skate App",
		storePrivatekeyCmd(),
		createPrivatekeyCmd(),
    retrievePrivateKeyCmd(),
		buildinfo.BuildInfoCmd(VERSION),
	)
}
