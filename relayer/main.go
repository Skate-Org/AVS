package main

// NOTE:
import (
	"context"

	libcmd "github.com/Skate-Org/AVS/lib/cmd"
	"github.com/Skate-Org/AVS/lib/logging"
	clicmd "github.com/Skate-Org/AVS/relayer/cmd"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/spf13/viper"
)

func main() {
	cmd := clicmd.New()

	fig := figure.NewFigure("Skate RELAYER", "", true)
	cmd.SetHelpTemplate(fig.String() + "\n" + cmd.HelpTemplate())

	libcmd.SilenceErrUsage(cmd)

	// Create a new Logger instance
	logger := logging.NewLoggerWithConsoleWriter()
	logger.Info(viper.GetString("skate_app"))

	err := cmd.ExecuteContext(context.Background())
	if err != nil {
		logger.Fatal("Fatal error occurred!", "error", err.Error())
	}
}
