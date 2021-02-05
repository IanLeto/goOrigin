package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery()) // 防止panic
	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "incorrect API route")
	})
	pingGroup := g.Group("/ping")
	{
		pingGroup.GET("", func(context *gin.Context) {
			return
		})
	}
	//registerProject := g.Group("/reg")
	//{
	//	registerProject.Group("/reg-pro/", nil)
	//}
	//recordProject := g.Group("/record")
	//{
	//	recordProject.GET("/", nil)
	//	recordProject.PUT("/create", nil)
	//	recordProject.POST("/", nil)
	//}
	return g
}
