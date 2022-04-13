package main

import (
	"goOrigin/cmd"
	"os"
)

func main() {
	switch os.Getenv("mode") {
	case "debug":
	default:
		cmd.PreRun("./config.yaml")
		cmd.DebugServer()
	}

}
