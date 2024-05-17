package db

import (
	"os"

	"github.com/Skate-Org/AVS/lib/logging"
)

var (
	DbDir     = "data/relayer"
)

func init() {
	logger := logging.NewLoggerWithConsoleWriter()
	if err := os.MkdirAll(DbDir, os.ModePerm); err != nil {
		logger.Fatal("Create data directory false, please try manually", "data dir", DbDir)
		panic(err)
	}
}
