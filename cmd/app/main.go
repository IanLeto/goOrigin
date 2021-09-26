package main

import (
	"github.com/gin-gonic/gin"
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
	mode := PreRun()
	switch mode {
	default:
		DebugServer()
	}
}
