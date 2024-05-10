package db

import (
	"io"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/Skate-Org/AVS/lib/logging"
)

func NewFileLogger(DbDir string, fileName string) *logging.Logger {
	logFile := createLogFile(DbDir, fileName)
	plainWriter := zerolog.ConsoleWriter{Out: logFile, TimeFormat: time.RFC3339, NoColor: true}
	return logging.NewLogger(plainWriter)
}

func createLogFile(DbDir string, fileName string) io.Writer {
	logger := logging.NewLoggerWithConsoleWriter()
	logFilePath := filepath.Join(DbDir, fileName)
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		if err := os.MkdirAll(DbDir, os.ModePerm); err != nil {
			logger.Error("Create data directory false, please try manually", "data dir", DbDir)
			panic(err)
		}

		// Create the log file with sufficient permissions
		file, err := os.Create(logFilePath)
		if err != nil {
			panic(err)
		}
		return file
	}

	// Open the existing log file with append mode and sufficient permissions
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	return file
}
