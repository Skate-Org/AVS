package exec

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/Skate-Org/AVS/lib/logging"
)

var logger = logging.NewLoggerWithConsoleWriter()

func ExecBin(timeoutSecond time.Duration, binary string, args ...string) error {
	// Setup the command and its arguments
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecond*time.Second)
	defer cancel() // Ensure the context is canceled to free resources

	cmd := exec.CommandContext(ctx, binary, args...)

	// Getting stdout pipe, which will be connected to the command's standard output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.Error("Error obtaining stdout", "error", err)
		return err
	}

	// Start the command before reading from the pipe
	if err := cmd.Start(); err != nil {
		logger.Error("Error starting command", "error", err)
		return err
	}

	// Create a new reader to read from the stdout
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break // End of file is a natural closure of the stream
		}
		if err != nil {
			logger.Error("Error reading stdout", "error", err)
			break
		}
		// Print the line to the standard output of the current Go process
		fmt.Print(line)
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		logger.Error("Binary execution failed failed!", "error", err)
		return err
	}

	return nil
}
