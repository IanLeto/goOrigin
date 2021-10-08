package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cmd "goOrigin/cmd"
	"goOrigin/config"
	_ "goOrigin/docs"
	"goOrigin/internal/router"
	"goOrigin/pkg/utils"
	"net/http"
)

func DebugServer() {
	g := gin.New()
	router.Load(g, nil)
	utils.NoError(http.ListenAndServe(fmt.Sprintf("%s", config.Conf.Url), g))
}

func main() {
	mode := cmd.PreRun()
	switch mode {
	default:
		DebugServer()
	}
}
