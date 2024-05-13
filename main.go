package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func main() {
    // Replace "./binary" with the actual path to your binary
    binary := "./bin/operator"

    // Setup the command and its arguments
    cmd := exec.Command(binary, "monitor", "--signer-config", "1", "--verbose=false")

    // Getting stdout pipe, which will be connected to the command's standard output
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Printf("Error obtaining stdout: %s\n", err)
        return
    }

    // Start the command before reading from the pipe
    if err := cmd.Start(); err != nil {
        fmt.Printf("Error starting command: %s\n", err)
        return
    }

    // Create a new reader to read from the stdout
    reader := bufio.NewReader(stdout)
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            break  // End of file is a natural closure of the stream
        }
        if err != nil {
            fmt.Printf("Error reading stdout: %s\n", err)
            break
        }
        // Print the line to the standard output of the current Go process
        fmt.Print(line)
    }

    // Wait for the command to finish
    err = cmd.Wait()
    if err != nil {
        fmt.Printf("Command finished with error: %s\n", err)
    }
}

