package router

import (
	"github.com/DeanThompson/ginpprof"
	_ "github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"goOrigin/internal/router/cmdHandlers"
	"goOrigin/internal/router/indexHandlers"
	"goOrigin/internal/router/k8sHandlers"
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
		recordGroup.POST("/", recordHandlers.CreateIanRecord)
		recordGroup.DELETE("/", recordHandlers.DeleteIanRecord)
		recordGroup.PUT("/", recordHandlers.UpdateIanRecord)
		recordGroup.GET("/", recordHandlers.SelectIanRecord)
		recordGroup.POST("/append", recordHandlers.AppendIanRecord)
		//recordGroup.POST("/", recordHandlers.CreateIanRecord)
		//recordGroup.POST("/", recordHandlers.CreateIanRecord)
	}
	k8sGroup := g.Group("v1/k8s/deploy")
	{

		k8sGroup.POST("", k8sHandlers.CreateDeploy)
		k8sGroup.GET("", k8sHandlers.ListDeploy)
		k8sGroup.DELETE("", k8sHandlers.DeleteDeploy)
		k8sGroup.PUT("", k8sHandlers.UpdateDeploy)

	}
	k8sConfigGroup := g.Group("v1/k8s/configmap")
	{
		k8sConfigGroup.POST("")
	}

	k8sV2Group := g.Group("v2/k8s/deploy")
	{
		k8sV2Group.POST("", k8sHandlers.CreateV2Deploy)
	}

	cmdGroup := g.Group("/v1/cmd")
	{
		cmdGroup.POST("ping", cmdHandlers.Ping)
	}
	return g
}
