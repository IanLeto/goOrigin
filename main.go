package main

import (
	"goOrigin/cmd"
	"os"
)

func main() {
	// 该env 由k8s 生成
	switch os.Getenv("mode") {
	default:
		cmd.PreRun("./config.yaml")
		cmd.DebugServer()
	}

}
