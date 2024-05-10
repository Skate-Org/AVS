package db

import (
	"os"

	"github.com/Skate-Org/AVS/lib/logging"
)

var (
	DbDir  = "data/operator"
	logger = logging.NewLoggerWithConsoleWriter()
)

func init() {
	if err := os.MkdirAll(DbDir, os.ModePerm); err != nil {
		logger.Error("Create data directory false, please try manually", "data dir", DbDir)
		panic(err)
	}
}
