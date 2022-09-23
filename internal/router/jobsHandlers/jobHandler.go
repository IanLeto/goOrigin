package jobsHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"goOrigin/internal/params"
	"goOrigin/internal/service"
	"goOrigin/pkg/logger"
)

// CreateJob  @Summary
// @Description 创建主任务
// @Tags Ian
// @Accept json
// @param job body params.CreateJobRequest true "1"
// @param res body params.BaseResponseInfo true "1"
// @Router /v1/record [POST]
func CreateJob(c *gin.Context) {

	var (
		req = params.CreateJobRequest{}
		id  uint
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	id, err = service.CreateJob(c, req)
	params.BuildResponse(c, params.BuildInfo(id))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

}

// UpdateJob   @Summary
// @Description 创建主任务
// @Tags Ian
// @Accept json
// @param job body params.CreateJobRequest true "1"
// @param res body params.BaseResponseInfo true "1"
// @Router /v1/record [POST]
func UpdateJob(c *gin.Context) {

	var (
		req = &params.UpdateJobRequest{}
		id  uint
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		//logger.Logger.Errorf("fail to json %s", err)
		logger.Logger.Error(fmt.Sprintf("errors : %s", err))
		goto ERR
	}
	id, err = service.UpdateJob(c, req)
	if err != nil {
		//logger.Logger.Errorf("fail to ")
		logger.Logger.Error("errors")
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(id))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

// DeleteJob  @Summary
// @Description 创建主任务
// @Tags Ian
// @Accept json
// @param job body params.CreateJobRequest true "1"
// @param res body params.BaseResponseInfo true "1"
// @Router /v1/record [POST]
func DeleteJob(c *gin.Context) {
	var (
		//id  = utils.QueryInt(c, "id", 0)
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

// GetJobDetail  @Summary
// @Description 创建主任务
// @Tags Ian
// @Accept json
// @param job body params.CreateJobRequest true "1"
// @param res body params.BaseResponseInfo true "1"
// @Router /v1/record [POST]
func GetJobDetail(c *gin.Context) {
	var (
		//id  = utils.QueryInt(c, "id", 0)
		id  = cast.ToInt(c.Param("id"))
		err error
		res = &params.GetJobResponse{}
	)
	res, err = service.GetJobDetail(c, id)
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

}

// GetJobs  @Summary
// @Description 创建主任务
// @Tags Ian
// @Accept json
// @param job body params.CreateJobRequest true "1"
// @param res body params.BaseResponseInfo true "1"
// @Router /v1/record [POST]
func GetJobs(c *gin.Context) {
	var (
		//id  = utils.QueryInt(c, "id", 0)
		//id  = cast.ToInt(c.Param("id"))
		err error
		res = &params.GetJobsResponse{}
	)
	res, err = service.GetJobs(c)
	if err != nil {
		goto ERR
	}
	params.BuildResponse(c, params.BuildInfo(res))
	return
ERR:
	params.BuildErrResponse(c, params.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

// ExecJob  @Summary
// @Description 创建主任务
// @Tags Ian
// @Accept json
// @param job body params.CreateJobRequest true "1"
// @param res body params.BaseResponseInfo true "1"
// @Router /v1/record [POST]
func ExecJob(c *gin.Context) {

}
