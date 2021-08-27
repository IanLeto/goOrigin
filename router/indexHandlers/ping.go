package indexHandlers

import (
	"github.com/gin-gonic/gin"
	"goOrigin/router/baseHandlers"
	"net/http"
)

func Ping(c *gin.Context) {
	c.String(200, "format")
}

func NoRouterHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "incorrect API address")
}

func BaseInformationHandler(c *gin.Context) {
	baseHandlers.RenderResponse(c, map[string]string{
		"Version":    "0.0.1",
		"Maintainer": "ian.liu",
		"DocUrl":     "",
	}, nil)
}
