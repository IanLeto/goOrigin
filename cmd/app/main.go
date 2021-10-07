package main

import (
	"github.com/gin-gonic/gin"
	_ "goOrigin/cmd/app/docs"
	"goOrigin/internal/router"
	"goOrigin/pkg/utils"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com
// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email me@razeen.me
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @BasePath /api/v1

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
