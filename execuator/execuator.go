// Package execuator provides functions for running shell commands
// and capturing their output.
// This also supports streaming the results while it is executing.
package execuator

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// Run : Run a function and return the std output and error
// Accepts two inputs, the fully qualified path of a binary (string) and concatinated list of flags (string)
// Executes the binary with the given flags and
// returns stdOut (string) and stdErr (string)
// func Run( string,  string) (string, string)
func Run(binaryPath string, params string) (string, string) {
	fmt.Println("[+++++ xcap.libs.execuator.Run] Start Run for ", binaryPath, params)

	cmd := exec.Command(binaryPath, params)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	fmt.Println("[+++++ xcap.libs.execuator.Run] End Run for ", binaryPath, params)
	return outStr, errStr

}

// Shellout : Run and stream the output to std display, also retuen the std out and err
// Accepts two inputs, the fully qualified path of a binary (string) and concatinated list of flags (string)
// Executes the binary with the given flags and
// returns stdOut (string) and stdErr (string)
// func Shellout( string,  string) (string, string)
func Shellout(binaryPath string, params string) (string, string) {
	fmt.Println("[+++++ xcap.libs.execuator.Shellout] Start Shellout for ", binaryPath, params)

	cmd := exec.Command(binaryPath, params)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())

	fmt.Println("[+++++ xcap.libs.execuator.Shellout] End Shellout for ", binaryPath, params)
	return outStr, errStr
}
