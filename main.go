package main

import (
	"goOrigin/cmd"
	_ "goOrigin/docs"
	_ "goOrigin/pkg/moniter"
	"goOrigin/pkg/utils"
)

func main() {
	utils.NoError(cmd.RootCmd.Execute())
}
