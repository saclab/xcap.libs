// Package executor provides functions for running shell commands
// and capturing their output.
// This also supports streaming the results while it is executing.
package executor

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/saclab/xcap.libs/logger"
)

// Run : Run a function and return the std output and error
// Accepts two inputs, the fully qualified path of a binary (string) and concatinated list of flags (string)
// Executes the binary with the given flags and
// returns stdOut (string) and stdErr (string)
// func Run( string,  string) (string, string)
// NOTE: Maintains a 'drop in replacement' relationship with function Shellout
func Run(binaryPath string, params string, extra map[string]string) (string, string) {
	logger.InitD()
	logger.Info.Println("(xcap.libs.executor.Run) Start Run for ", binaryPath, params)

	cmd := exec.Command(binaryPath, params)
	if val, ok := extra["dir"]; ok {
		cmd.Dir = val
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		logger.Error.Println("(xcap.libs.executor.Run) cmd.Run() failed with : ", err)
	}

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	logger.Info.Println("(xcap.libs.executor.Run)  End Run for ", binaryPath, params)
	return outStr, errStr
}

// Shellout : Run and stream the output to std display, also retuen the std out and err
// Accepts two inputs, the fully qualified path of a binary (string) and concatinated list of flags (string)
// Executes the binary with the given flags and
// returns stdOut (string) and stdErr (string)
// func Shellout( string,  string) (string, string)
// NOTE: Maintains a 'drop in replacement' relationship with function Run
func Shellout(binaryPath string, params string, extra map[string]string) (string, string) {
	logger.InitD()
	logger.Info.Println("(xcap.libs.executor.Shellout) Start Shellout for ", binaryPath, params)

	cmd := exec.Command(binaryPath, params)
	if val, ok := extra["dir"]; ok {
		cmd.Dir = val
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		logger.Error.Println("(xcap.libs.executor.Shellout) cmd.Run() failed with : ", err)
	}

	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())

	logger.Info.Println("(xcap.libs.executor.Shellout)  End Shellout for ", binaryPath, params)
	return outStr, errStr
}
