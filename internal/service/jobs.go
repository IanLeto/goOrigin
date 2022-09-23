package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
	"goOrigin/pkg/clients"
	"goOrigin/pkg/logger"
)

func CreateJob(c *gin.Context, req params.CreateJobRequest) (uint, error) {
	var (
		job = model.Job{
			Targets:    req.Targets,
			FilePath:   req.FilePath,
			Name:       req.Name,
			Type:       req.Type,
			StrategyID: req.StrategyID,
		}
		log    = logger.NewLogger()
		err    error
		client *elastic.Client
	)
	client, err = clients.NewESClient()
	if err != nil {
		log.Error(fmt.Sprintf("创建es client 失败 %s", err.Error()))
		goto ERR
	}

	for _, id := range req.ScriptIDs {
		client.Search().Index("script").Query(elastic.NewBoolQuery().Filter(elastic.NewTermQuery("ID", id)))
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

func UpdateJob(c *gin.Context, req *params.UpdateJobRequest) (uint, error) {
	var (
		job = model.Job{
			ID: req.ID,
			//Target:   req.Target,
			FilePath: req.FilePath,
			Name:     req.Name,
		}
		err error
	)
	if err != nil {
		logrus.Errorf("更新 job 失败 %s", err)
		goto ERR
	}
	err = job.Update()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("errors : %s", err))
		goto ERR
	}
	return job.ID, err
ERR:
	{
		return uint(0), nil
	}
}

func DeleteJob(c *gin.Context, id int) error {
	var (
		err error
		job = &model.Job{
			ID: uint(id),
		}
	)
	err = job.Delete()
	return err
}
func GetJobDetail(c *gin.Context, id int) (*params.GetJobResponse, error) {
	var (
		err error
		job = &model.Job{ID: uint(id)}
	)
	tJob, err := job.QueryDetail()
	if err != nil {
		return nil, err
	}
	response := &params.GetJobResponse{
		ID:       tJob.ID,
		Name:     tJob.Name,
		FilePath: tJob.FilePath,
	}

	return response, err

}
