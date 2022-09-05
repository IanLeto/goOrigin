package promHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goOrigin/internal/params"
	"goOrigin/internal/service"
)

// QueryWeight  @Summary
// @Description 创建主任务
// @Tags Ian
// @Accept json
// @param job body params.QueryWeightRequest true "1"
// @param res body params.BaseResponseInfo true "1"
// @Router /v1/record [POST]
func QueryWeight(c *gin.Context) {
	var (
		req = params.QueryWeightRequest{}
		err error
		res string
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		goto ERR
	}
	res, err = service.QueryPromWeight(c, &req)
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

}
