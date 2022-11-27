package main

import (
	"goOrigin/cmd"
	_ "goOrigin/docs"
	_ "goOrigin/pkg/moniter"
)

func main() {

	cmd.Execute()
	// 该env 由k8s 生成
	//switch os.Getenv("mode") {
	//default:
	//	cmd.PreRun("./config.yaml")
	//	cmd.DebugServer()
	//}

}
