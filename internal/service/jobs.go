package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
)

func CreateJob(c *gin.Context, req params.CreateJobRequest) (uint, error) {
	var (
		job = model.NewJob(req)
		err error
	)
	if err != nil {
		logrus.Errorf("创建 job 失败 %s", err)
		goto ERR
	}
	err = job.Create()
	if err != nil {
		goto ERR
	}
	return job.ID, err
ERR:
	{
		return uint(0), nil
	}
}
