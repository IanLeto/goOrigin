package scriptHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"goOrigin/internal/params"
	"goOrigin/internal/service"
)

func CreateScript(c *gin.Context) {
	var (
		req = params.CreateScriptRequest{}
		res interface{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	res, err = service.CreateScript(c, req)
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func DeleteScript(c *gin.Context) {
	var (
		id  = cast.ToInt(c.Param("id"))
		err error
	)
	err = service.DeleteJob(c, id)
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(id))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func QueryScriptList(c *gin.Context) {
	var (
		req = params.QueryScriptRequest{
			Type: c.Query("type"),
			Key:  c.Query("key"),
			Tags: c.Query("tags"),
		}
		res = &params.QueryScriptListResponse{}
		err error
	)

	res, err = service.QueryScript(c, req)
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func RunScript(c *gin.Context) {
	var (
		err error
		id  = c.Query("id")
	)

	res, err := service.RunScript(c, id)
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
