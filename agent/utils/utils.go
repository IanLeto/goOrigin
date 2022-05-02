package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func RunShell(content string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	var err error
	var errStdout, errStderr error
	cmd := exec.Command("ls", "-lah")
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StdoutPipe()
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err = cmd.Start()
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("%s", err)
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Println(outStr, errStr)
}
