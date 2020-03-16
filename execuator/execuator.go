package execuator

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// HelloWorld : Just a test function
func HelloWorld() {
	fmt.Println("Hello from execuator package")

}

// Run : Run a function
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

// Shellout : Run and stream the output
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
	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	fmt.Println("[+++++ xcap.libs.execuator.Shellout] End Shellout for ", binaryPath, params)
	return outStr, errStr
}
