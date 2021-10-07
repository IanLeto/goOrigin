package main

import (
	"github.com/gin-gonic/gin"
	cmd "goOrigin/cmd"
	"goOrigin/internal/router"
	"goOrigin/pkg/utils"
	"net/http"
)

func DebugServer() {
	g := gin.New()
	router.Load(g, nil)
	utils.NoError(http.ListenAndServe("localhost:8008", g))
}

func main() {
	mode := cmd.PreRun()
	switch mode {
	default:
		DebugServer()
	}
}
