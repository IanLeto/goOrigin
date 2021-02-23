package router

import (
	"github.com/gin-gonic/gin"
	"goOrigin/router/indexHandler"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery()) // 防止panic

	g.NoRoute(indexHandler.NoRouterHandler)
	indexGroup := g.Group("/")
	{
		indexGroup.GET("ping", indexHandler.Ping)
	}

	return g
}
