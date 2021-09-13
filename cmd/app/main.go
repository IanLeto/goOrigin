package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goOrigin/internal/router"
	"goOrigin/pkg/utils"
	"net/http"
)

func DebugServer() {
	g := gin.New()
	router.Load(g, nil)
	utils.NoError(http.ListenAndServe(":"+utils.ConvOrDefaultString(viper.Get("port"), "8080"), g))
}

func main() {
	mode := PreRun()
	switch mode {
	default:
		DebugServer()
	}
}
