package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model"
	"goOrigin/pkg/logger"
	"goOrigin/pkg/storage"
	"strings"
)

func CreateJob(c *gin.Context, req V1.CreateJobRequest) (uint, error) {
	var (
		job = model.Job{
			Targets:    req.Targets,
			FilePath:   req.FilePath,
			Name:       req.Name,
			Type:       req.Type,
			StrategyID: req.StrategyID,
			ScriptIDS:  req.ScriptIDs,
		}
		log = logger.NewLogger()
		err error
		//client   *elastic.Client
		//playbook = map[string]string{}
	)
	//scripts, err := model.BoolQueryScript(c, client)
	//client, err = clients.NewESClient()
	if err != nil {
		log.Error(fmt.Sprintf("创建es client 失败 %s", err.Error()))
		goto ERR
	}
	//res, err = client.Search().Index("script").Query(elastic.NewBoolQuery().Filter(elastic.NewTermQuery("ID", req.ScriptIDs))).Do(c)
	//if err != nil {
	//	log.Error(fmt.Sprintf("查询es client 失败 %s", err.Error()))
	//	goto ERR
	//}

	//for _, id := range req.ScriptIDs {
	//	for i := 0; i < len(scripts); i++ {
	//		if id == scripts[i].ID {
	//			playbook[id] = scripts[i].Content
	//			scripts[i] = scripts[len(scripts)-1]
	//			scripts[len(scripts)-1] = nil
	//			scripts = scripts[:len(scripts)-1]
	//		}
	//	}
	//}
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

func UpdateJob(c *gin.Context, req *V1.UpdateJobRequest) (uint, error) {
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
func GetJobDetail(c *gin.Context, id int) (*V1.GetJobResponse, error) {
	var (
		err error
		job = &model.Job{ID: uint(id)}
	)
	tJob, err := job.QueryDetail()
	if err != nil {
		return nil, err
	}
	response := &V1.GetJobResponse{
		ID:       tJob.ID,
		Name:     tJob.Name,
		FilePath: tJob.FilePath,
		Target:   strings.Split(tJob.Target, ","),
	}

	return response, err

}

func GetJobs(c *gin.Context) (*V1.GetJobsResponse, error) {
	var (
		err   error
		tJobs []*mysql.TJob
		infos []*V1.GetJobsResponseInfo
	)

	err = storage.GlobalMySQL.Table("t_jobs").Find(&tJobs).Error
	if err != nil {
		return nil, err
	}
	for _, tJob := range tJobs {
		infos = append(infos, &V1.GetJobsResponseInfo{
			ID:        tJob.ID,
			Name:      tJob.Name,
			Content:   "",
			ScriptIDs: strings.Split(tJob.ScriptIDs, ","),
		})
	}
	return &V1.GetJobsResponse{
		Infos: infos,
	}, err

}

func RunJob(c *gin.Context, req *V1.RunJobRequest) (*V1.RunJobResponse, error) {
	var (
		err  error
		job  = &model.Job{}
		tJob = &mysql.TJob{}
		//infos []*params.GetJobsResponseInfo
	)
	if err != nil {
		goto ERR
	}

	err = storage.GlobalMySQL.Table("t_jobs").Find(&tJob).Error
	if err != nil {
		return nil, err
	}
	job = &model.Job{
		ID:        tJob.ID,
		Targets:   strings.Split(tJob.ScriptIDs, ","),
		FilePath:  tJob.FilePath,
		Name:      tJob.Name,
		Type:      tJob.Type,
		ScriptIDS: strings.Split(tJob.ScriptIDs, ","),
	}
	go func() {

		_ = job.Exec(c)
	}()
	return &V1.RunJobResponse{}, err
ERR:
	{
		return nil, err
	}
}
