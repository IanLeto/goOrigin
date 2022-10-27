package indexHandlers

import (
	"github.com/gin-gonic/gin"
	"goOrigin/internal/router/baseHandlers"
	"net/http"
)

func Ping(c *gin.Context) {

	baseHandlers.RenderData(c, map[string]string{
		"Version":    "0.0.1",
		"Maintainer": "ian.liu",
		"DocUrl":     "",
	}, nil)
}

func Prom(handler http.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		handler.ServeHTTP(context.Writer, context.Request)
	}
}

func NoRouterHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "incorrect API address")
}
