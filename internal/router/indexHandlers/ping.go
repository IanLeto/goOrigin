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

func NoRouterHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "incorrect API address")
}
