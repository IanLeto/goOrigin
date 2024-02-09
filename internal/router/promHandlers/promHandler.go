package promHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
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
		req = V1.QueryWeightRequest{}
		err error
		res string
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		goto ERR
	}
	res, err = logic.QueryPromWeight(c, &req)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

}
