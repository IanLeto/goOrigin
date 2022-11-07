package main

import (
	"goOrigin/cmd"
	_ "goOrigin/docs"
	_ "goOrigin/pkg/moniter"
)

func main() {
	cmd.Execute()
}
