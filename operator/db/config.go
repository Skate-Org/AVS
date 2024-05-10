package db

import (
	"os"

	"skatechain.org/lib/logging"
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
