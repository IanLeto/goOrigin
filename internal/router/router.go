package router

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goOrigin/internal/router/trans_type"
	"goOrigin/pkg/moniter"

	"goOrigin/internal/router/indexHandlers"
	"goOrigin/internal/router/recordHandlers"
	"goOrigin/internal/router/topoHandlers"
	_ "goOrigin/pkg/collector"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	//ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	//
	//otelShutdown, err := clients.SetupOTelSDK(ctx, func() map[string]string {
	//	traceInfo := map[string]string{}
	//	t := reflect.TypeOf(config.ConfV2.Trace)
	//	v := reflect.ValueOf(config.ConfV2.Trace)
	//	for i := 0; i < t.NumField(); i++ {
	//		traceInfo[t.Field(i).Tag.Get("json")] = v.Field(i).String()
	//	}
	//	return traceInfo
	//}())
	//utils.NoError(err)
	//// Handle shutdown properly so nothing leaks.
	//defer func() {
	//	if err != nil {
	//		_ = otelShutdown(ctx)
	//	}
	//}()
	g.Use(CORSMiddleware())
	//g.Use(AuthMiddleware())
	g.Use(gin.Recovery()) // 防止panic
	g.NoRoute(indexHandlers.NoRouterHandler)
	//g.Use(Jaeger())
	pprof.Register(g, "debug/pprof")
	//g.Use(otelgin.Middleware("my-service"))
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	indexGroup := g.Group("/")
	{
		indexGroup.GET("ping", indexHandlers.Ping)
		indexGroup.GET("config", indexHandlers.ConfigInfo)
		//indexGroup.GET("podInfo", indexHandlers.ConfigCheck)
		indexGroup.GET("userInfo", indexHandlers.GetUser)
		indexGroup.GET("", indexHandlers.HttpProxy)
		//indexGroup.GET("proxy", indexHandlers.HttpProxy)
		indexGroup.GET("metrics", indexHandlers.Prom(promhttp.Handler()))

	}
	execGroup := g.Group("/v1/exec")
	{
		execGroup.POST("/:id")
	}
	recordGroup := g.Group("/v1/record")
	{
		recordGroup.POST("", recordHandlers.CreateRecord)
		recordGroup.GET("", recordHandlers.QueryRecords)
		recordGroup.PUT("", recordHandlers.UpdateRecord)
		recordGroup.DELETE("", recordHandlers.DeleteRecord)
		recordGroup.POST("/file", recordHandlers.CreateFileRecord)

	}
	ginpprof.WrapGroup(recordGroup)

	nodev2Group := g.Group("v2/node")
	{
		nodev2Group.POST("", topoHandlers.CreateNode)
		//nodev2Group.POST("/batch", topoHandlers.CreateNodes)
		nodev2Group.PUT("", topoHandlers.UpdateNode)
		nodev2Group.DELETE("", topoHandlers.DeleteNode)
		nodev2Group.GET("", topoHandlers.GetNodeDetail)
		nodev2Group.GET("/search", topoHandlers.GetNodeDetail)

		// ✅ 加上这一行
		nodev2Group.GET("/list", topoHandlers.ListNodes)
	}

	trasnGroup := g.Group("v1/trans")
	{
		trasnGroup.POST("", trans_type.CreateTransInfo)
		trasnGroup.POST("list", trans_type.GetTransInfoList)
		trasnGroup.DELETE("", trans_type.DeleteTransInfo)
		trasnGroup.PATCH("", trans_type.UpdateTransInfo)
	}

	odav2Group := g.Group("v2/pub")
	{
		odav2Group.GET("list", nil)
	}
	metricGroup := g.Group("v1/metric")
	{
		ops := promhttp.HandlerOpts{
			EnableOpenMetrics: false,
		}
		metricGroup.GET("", gin.WrapH(promhttp.HandlerFor(moniter.Reg, ops)))

	}
	configGroup := g.Group("v1/config")
	{
		configGroup.GET("", indexHandlers.ConfigInfo)
	}

	return g
}
