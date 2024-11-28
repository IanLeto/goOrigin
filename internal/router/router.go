package router

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"goOrigin/pkg/moniter"

	//"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"goOrigin/config"
	"goOrigin/internal/router/cmdHandlers"
	"goOrigin/internal/router/indexHandlers"
	"goOrigin/internal/router/jobsHandlers"
	"goOrigin/internal/router/k8sHandlers"
	"goOrigin/internal/router/promHandlers"
	"goOrigin/internal/router/recordHandlers"
	"goOrigin/internal/router/scriptHandlers"
	"goOrigin/internal/router/topoHandlers"
	_ "goOrigin/pkg/collector"
	"io"
	"net/http"
)

func newTracer(svc, collectorEndpoint string) (opentracing.Tracer, io.Closer) {
	cfg := jaegerConfig.Configuration{

		ServiceName: svc,
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const",
			Param: 1, // 1 全采 0 不采
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: config.Conf.Backend.JaegerConfig.Address,
		},
	}
	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
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
		indexGroup.GET("configCheck", indexHandlers.ConfigCheck)
		indexGroup.GET("podInfo", indexHandlers.ConfigCheck)
		indexGroup.GET("userInfo", indexHandlers.GetUser)
		indexGroup.GET("", indexHandlers.HttpProxy)
		//indexGroup.GET("proxy", indexHandlers.HttpProxy)
		indexGroup.GET("metrics", indexHandlers.Prom(promhttp.Handler()))
		indexGroup.GET("jobInfo", indexHandlers.JobInfo)

	}
	execGroup := g.Group("/v1/exec")
	{
		execGroup.POST("/:id")
	}
	recordGroup := g.Group("/v1/record")
	{
		recordGroup.POST("", recordHandlers.CreateRecord)
		recordGroup.GET("", recordHandlers.QueryRecord)
		recordGroup.PUT("", recordHandlers.UpdateRecord)
		recordGroup.DELETE("", recordHandlers.DeleteIanRecord)
		recordGroup.POST("/append", recordHandlers.AppendIanRecord)

	}
	ginpprof.WrapGroup(recordGroup)
	k8sGroup := g.Group("v1/k8s/deploy")
	{
		k8sGroup.POST("", k8sHandlers.CreateDeploy)
		k8sGroup.GET("", k8sHandlers.ListDeploy)
		k8sGroup.DELETE("", k8sHandlers.DeleteDeploy)
		k8sGroup.PUT("", k8sHandlers.UpdateDeploy)
		k8sGroup.POST("dynamic", k8sHandlers.CreateDeployDynamic)
		k8sGroup.DELETE("dynamic", k8sHandlers.DeleteDeployDynamic)
		k8sGroup.PUT("dynamic", k8sHandlers.UpdateDeployDynamic)
		k8sGroup.GET("log", k8sHandlers.GetCurrentLogs)
		k8sGroup.GET("pod", k8sHandlers.GetPods)

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

	jobGroup := g.Group("/v1/job")
	{
		jobGroup.POST("/", jobsHandlers.CreateJob)
		jobGroup.DELETE(":id", jobsHandlers.DeleteJob)
		jobGroup.PUT("/", jobsHandlers.UpdateJob)
		jobGroup.GET(":id", jobsHandlers.GetJobDetail)
		jobGroup.GET("", jobsHandlers.GetJobs)
		jobGroup.POST("run/:id", jobsHandlers.RunJob)
	}
	// 远程prom数据展示
	tencentGroup := g.Group("/v1/dashboard")
	{
		// 基础数据
		tencentGroup.POST(":id", jobsHandlers.GetJobDetail)
	}

	scriptGroup := g.Group("v1/script")
	{
		// 基础数据
		scriptGroup.GET("", scriptHandlers.QueryScriptList)
		scriptGroup.GET("exec", scriptHandlers.RunScript)
		scriptGroup.POST("", scriptHandlers.CreateScript)
		scriptGroup.DELETE("", scriptHandlers.DeleteScript)

	}
	promGroup := g.Group("v1/prom")
	{
		promGroup.POST("weight", promHandlers.QueryWeight)

	}

	nodev2Group := g.Group("v2/node")
	{
		nodev2Group.POST("", topoHandlers.CreateNode)
		nodev2Group.POST("/batch", topoHandlers.CreateNodes)
		nodev2Group.PUT("", topoHandlers.UpdateNode)
		nodev2Group.DELETE("", topoHandlers.DeleteNode)
		nodev2Group.POST("/list", topoHandlers.GetNodes)
		nodev2Group.GET("", topoHandlers.GetNodeDetail)
		nodev2Group.GET("/search", topoHandlers.GetNodeDetail)
	}

	topov2Group := g.Group("v2/topo")
	{
		topov2Group.GET("list", topoHandlers.GetTopoList)
	}
	metricGroup := g.Group("v1/metric")
	{
		ops := promhttp.HandlerOpts{
			EnableOpenMetrics: false,
		}
		metricGroup.GET("", gin.WrapH(promhttp.HandlerFor(moniter.Reg, ops)))
	}
	return g
}
