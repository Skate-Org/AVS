package main

// NOTE:
import (
	"context"

	clicmd "github.com/Skate-Org/AVS/kms/cmd"
	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/logging"

	figure "github.com/common-nighthawk/go-figure"
)

func main() {
	cmd := clicmd.New()

	fig := figure.NewFigure("Key Management Service", "", true)
	cmd.SetHelpTemplate(fig.String() + "\n" + cmd.HelpTemplate())

	libcmd.SilenceErrUsage(cmd)

	// Create a new Logger instance
	logger := logging.NewLoggerWithConsoleWriter()

	err := cmd.ExecuteContext(context.Background())
	if err != nil {
		logger.Fatal("Fatal error occurred!", "error", err.Error())
	}
}
