package router

import (
	"github.com/gin-gonic/gin"
	"goOrigin/internal/router/indexHandlers"
	"goOrigin/internal/router/userHandlers"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery()) // 防止panic

	g.NoRoute(indexHandlers.NoRouterHandler)
	indexGroup := g.Group("/")
	{
		indexGroup.GET("ping", indexHandlers.Ping)

	}
	userGroup := g.Group("/v1/user")
	{
		userGroup.POST("", userHandlers.Create)
		userGroup.PUT("/:id", userHandlers.Update)
		userGroup.GET("", userHandlers.List)
		userGroup.GET("/:username", userHandlers.Get)
		userGroup.DELETE("/:id", userHandlers.Delete)
	}

	execGroup := g.Group("/v1/exec")
	{
		execGroup.POST("/:id")
	}
	return g
}
