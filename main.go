package main

import (
	"fmt"
	cmd "goOrigin/cmd"
)

func main() {
	//switch os.Getenv("mode") {
	//case "debug":
	fmt.Println(cmd.RootCmd.Execute())
	//default:
	//	cmd.PreRun("")
	//	cmd.DebugServer()
	//}

}
