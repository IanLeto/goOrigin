package main

import (
	"goOrigin/cmd"
	_ "goOrigin/docs"
	_ "goOrigin/pkg/moniter"
)

func main() {
	//fmt.Println(fmt.Sprintf(`"%s"`, "dss"))
	cmd.Execute()

}
