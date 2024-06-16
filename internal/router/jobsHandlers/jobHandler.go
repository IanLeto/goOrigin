package jobsHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
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
		req = V1.CreateJobRequest{}
		id  uint
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	id, err = logic.CreateJob(c, req)
	V1.BuildResponse(c, V1.BuildInfo(id))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

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
		req = &V1.UpdateJobRequest{}
		id  uint
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		//logger.Logger.Errorf("fail to json %s", err)
		logger.Logger.Error(fmt.Sprintf("errors : %s", err))
		goto ERR
	}
	id, err = logic.UpdateJob(c, req)
	if err != nil {
		//logger.Logger.Errorf("fail to ")
		logger.Logger.Error("errors")
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(id))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
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
	err = logic.DeleteJob(c, id)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(id))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

}

// GetJobDetail  @Summary
// @Description 创建主任务
// @Tags Ian
// @Accept json
// @param job body params.CreateJobRequest true "1"
// @param res body params.BaseResponseInfo true "1"
// @Router /v1/record [POST]
func GetJobDetail(c *gin.Context) {
	//	var (
	//		//id  = utils.QueryInt(c, "id", 0)
	//		id  = cast.ToInt(c.Param("id"))
	//		err error
	//		res = &V1.GetJobResponse{}
	//	)
	//	res, err = logic.GetJobDetail(c, id)
	//	if err != nil {
	//		goto ERR
	//	}
	//	V1.BuildResponse(c, V1.BuildInfo(res))
	//	return
	//ERR:
	//	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

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
		res = &V1.GetJobsResponse{}
	)
	res, err = logic.GetJobs(c)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func RunJob(c *gin.Context) {

	var (
		req = &V1.RunJobRequest{}
		res = &V1.RunJobResponse{}
		err error
	)
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}
	res, err = logic.RunJob(c, req)
	if err != nil {
		goto ERR
	}
	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))

}
