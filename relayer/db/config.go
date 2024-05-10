package db

import (
	"os"

	"skatechain.org/lib/logging"
)

var (
	DbDir     = "data/relayer"
	Separator = []byte("::")
)

func init() {
	logger := logging.NewLoggerWithConsoleWriter()
	if err := os.MkdirAll(DbDir, os.ModePerm); err != nil {
		logger.Fatal("Create data directory false, please try manually", "data dir", DbDir)
		panic(err)
	}
}
