package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/db"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
	"goOrigin/pkg/logger"
	"goOrigin/pkg/storage"
	"strings"
)

func CreateJob(c *gin.Context, req params.CreateJobRequest) (uint, error) {
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

func GetJobs(c *gin.Context) (*params.GetJobsResponse, error) {
	var (
		err   error
		tJobs []*db.TJob
		infos []*params.GetJobsResponseInfo
	)

	err = storage.GlobalMySQL.Table("t_jobs").Find(&tJobs).Error
	if err != nil {
		return nil, err
	}
	for _, tJob := range tJobs {
		infos = append(infos, &params.GetJobsResponseInfo{
			ID:        tJob.ID,
			Name:      tJob.Name,
			Content:   "",
			ScriptIDs: strings.Split(tJob.ScriptIDs, ","),
		})
	}
	return &params.GetJobsResponse{
		Infos: infos,
	}, err

}
