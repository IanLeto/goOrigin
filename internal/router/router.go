package router

import (
	"github.com/DeanThompson/ginpprof"
	_ "github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"goOrigin/internal/router/indexHandlers"
	"goOrigin/internal/router/recordHandlers"
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
	ginpprof.WrapGroup(userGroup)

	execGroup := g.Group("/v1/exec")
	{
		execGroup.POST("/:id")
	}
	recordGroup := g.Group("/v1/record")
	{
		recordGroup.POST("/", recordHandlers.CreateRecord)
	}
	return g
}
