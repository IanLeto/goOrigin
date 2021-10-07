package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "goOrigin/cmd/app/docs"
	"goOrigin/internal/router/ianHandlers"
	"goOrigin/internal/router/indexHandlers"
	"goOrigin/internal/router/userHandlers"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery()) // 防止panic
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
	ianGroup := g.Group("/ian")
	{
		ianGroup.POST("/addForm", ianHandlers.AddDayForm)
		ianGroup.POST("/updateForm", ianHandlers.UpdateForm)
		ianGroup.POST("/SelectForm", ianHandlers.SelectForm)
	}
	execGroup := g.Group("/v1/exec")
	{
		execGroup.POST("/:id")
	}
	return g
}
