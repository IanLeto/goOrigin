package main

import (
	"goOrigin/cmd"
	"os"
)

func main() {
	switch os.Getenv("mode") {
	default:
		cmd.PreRun("./config.yaml")
		cmd.DebugServer()
	}

}
