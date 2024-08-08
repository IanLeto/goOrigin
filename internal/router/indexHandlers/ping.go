package indexHandlers

import (
	"encoding/json"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"goOrigin/config"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/router/baseHandlers"
	"goOrigin/pkg/utils"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Ping(c *gin.Context) {
	//file, _ := os.OpenFile("cpu", os.O_CREATE|os.O_RDWR, 0777)
	//err := pprof.StartCPUProfile(file)
	//if err != nil {
	//	panic(err)
	//}
	//defer pprof.StopCPUProfile()

	baseHandlers.RenderData(c, config.ConfV2, nil)
}

func ConfigInfo(c *gin.Context) {
	var (
		result = make(map[string]interface{})
	)
	span := trace.SpanFromContext(c.Request.Context())

	defer span.End()
	v, err := json.Marshal(config.ConfV2)
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

func HttpProxy(c *gin.Context) {
	var (
		loginUrl  string
		targetURL *url.URL
	)
	r := c.Request
	loginUrl, _ = conv.String(c.GetHeader("loginUrl"))
	targetURL, _ = url.Parse(utils.GetLoginUrlOrigin(loginUrl))

	// 保留请求路径，只修改目标 URL 的主机和协议部分
	r.URL.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	r.Host = targetURL.Host
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	proxy.ServeHTTP(c.Writer, r)
}

func GetUser(c *gin.Context) {
	var (
		token    string
		loginUrl string
		err      error
		user     entity.User
	)
	token = c.GetHeader("token")
	if token == "" {
		baseHandlers.RenderData(c, "no token", nil)
		return
	}
	loginUrl, err = conv.String(c.Value("loginUrl"))
	utils.NoError(err)
	userStr := entity.UserFromToken(token)
	user = &userStr
	u, ok := user.ToUserEntity(token, "", loginUrl).(*entity.CpaasUserEntity)
	if !ok {
		baseHandlers.RenderData(c, "error", nil)
		return
	}

	if err != nil {
		baseHandlers.RenderData(c, nil, err)
		return
	}

	baseHandlers.RenderData(c, u, err)
}

func Prom(handler http.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		handler.ServeHTTP(context.Writer, context.Request)
	}
}

func NoRouterHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "incorrect API address")
}
