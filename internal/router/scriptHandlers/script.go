package scriptHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
)

func CreateScript(c *gin.Context) {
	var (
		req = V1.CreateScriptRequest{}
		res interface{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	res, err = logic.CreateScript(c, req)
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func DeleteScript(c *gin.Context) {
	var (
		id  = cast.ToInt(c.Param("id"))
		err error
	)
	err = logic.DeleteJob(c, id)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(id))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func QueryScriptList(c *gin.Context) {
	var (
		req = V1.QueryScriptRequest{
			Type: c.Query("type"),
			Key:  c.Query("key"),
			Tags: c.Query("tags"),
		}
		res = &V1.QueryScriptListResponse{}
		err error
	)

	res, err = logic.QueryScript(c, req)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func RunScript(c *gin.Context) {
	var (
		err error
		id  = c.Query("id")
	)

	res, err := logic.RunScript(c, id)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}
