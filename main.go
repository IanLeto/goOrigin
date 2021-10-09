package main

import (
	cmd "goOrigin/cmd"
	"goOrigin/pkg/utils"
	"os"
)

func main() {
	switch os.Getenv("mode") {
	case "debug":
		utils.NoError(cmd.RootCmd.Execute())
	default:
		cmd.PreRun("")
		cmd.DebugServer()
	}

}
