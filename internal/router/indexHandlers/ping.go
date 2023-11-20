package indexHandlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goOrigin/config"
	"goOrigin/internal/router/baseHandlers"
	"net/http"
)

func Ping(c *gin.Context) {
	//file, _ := os.OpenFile("cpu", os.O_CREATE|os.O_RDWR, 0777)
	//err := pprof.StartCPUProfile(file)
	//if err != nil {
	//	panic(err)
	//}
	//defer pprof.StopCPUProfile()

	baseHandlers.RenderData(c, config.Conf, nil)
}

func ConfigInfo(c *gin.Context) {
	var (
		result = make(map[string]interface{})
	)
	v, err := json.Marshal(config.Conf)
	if err != nil {
		baseHandlers.RenderData(c, nil, err)
		return
	}

	err = json.Unmarshal(v, &result)
	baseHandlers.RenderData(c, result, err)
}

func ConfigCheck(c *gin.Context) {

	res, err := config.Conf.Backend.K8sConfig.Check()
	if err != nil {
		baseHandlers.RenderData(c, res, err)
		return
	}

	baseHandlers.RenderData(c, "ok", err)
}

func Prom(handler http.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		handler.ServeHTTP(context.Writer, context.Request)
	}
}

func NoRouterHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "incorrect API address")
}
