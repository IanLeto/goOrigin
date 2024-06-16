package logic

import (
	"github.com/gin-gonic/gin"
	"goOrigin/API/V1"
)

func CreateJob(c *gin.Context, req V1.CreateJobRequest) (uint, error) {
	panic("implement me")
	//	var (
	//		job = entity.Job{
	//			Targets:    req.Targets,
	//			FilePath:   req.FilePath,
	//			Name:       req.Name,
	//			Type:       req.Type,
	//			StrategyID: req.StrategyID,
	//			ScriptIDS:  req.ScriptIDs,
	//		}
	//		log = logger.NewLogger()
	//		err error
	//		//client   *elastic.Client
	//		//playbook = map[string]string{}
	//	)
	//	//scripts, err := model.BoolQueryScript(c, client)
	//	//client, err = clients.NewESClient()
	//	if err != nil {
	//		log.Error(fmt.Sprintf("创建es client 失败 %s", err.Error()))
	//		goto ERR
	//	}
	//	err = job.Create()
	//	if err != nil {
	//		goto ERR
	//	}
	//	return job.ID, err
	//ERR:
	//	{
	//		return uint(0), nil
	//	}
}

func UpdateJob(c *gin.Context, req *V1.UpdateJobRequest) (uint, error) {
	panic("implement me")
	//	var (
	//		job = entity.Job{
	//			ID: req.ID,
	//			//Target:   req.Target,
	//			FilePath: req.FilePath,
	//			Name:     req.Name,
	//		}
	//		err error
	//	)
	//	if err != nil {
	//		logrus.Errorf("更新 job 失败 %s", err)
	//		goto ERR
	//	}
	//	err = job.Update()
	//	if err != nil {
	//		logger.Logger.Error(fmt.Sprintf("errors : %s", err))
	//		goto ERR
	//	}
	//	return job.ID, err
	//ERR:
	//	{
	//		return uint(0), nil
	//	}
}

func DeleteJob(c *gin.Context, id int) error {
	panic("implement me")
	//	var (
	//		err error
	//		job = &entity.Job{
	//			ID: uint(id),
	//		}
	//	)
	//	err = job.Delete()
	//	return err
	//}
	//func GetJobDetail(c *gin.Context, id int) (*V1.GetJobResponse, error) {
	//	var (
	//		err error
	//		job = &entity.Job{ID: uint(id)}
	//	)
	//	tJob, err := job.QueryDetail()
	//	if err != nil {
	//		return nil, err
	//	}
	//	response := &V1.GetJobResponse{
	//		//ID:       tJob.ID,
	//		Name:     tJob.Name,
	//		FilePath: tJob.FilePath,
	//		Target:   strings.Split(tJob.Target, ","),
	//	}
	//
	//	return response, err

}

func GetJobs(c *gin.Context) (*V1.GetJobsResponse, error) {
	panic("implement me")
	//var (
	//	err   error
	//	tJobs []*dao.TJob
	//	infos []*V1.GetJobsResponseInfo
	//)
	//
	//err = storage.GlobalMySQL.Table("t_jobs").Find(&tJobs).Error
	//if err != nil {
	//	return nil, err
	//}
	//for _, tJob := range tJobs {
	//	infos = append(infos, &V1.GetJobsResponseInfo{
	//		//ID:        tJob.ID,
	//		Name:      tJob.Name,
	//		Content:   "",
	//		ScriptIDs: strings.Split(tJob.ScriptIDs, ","),
	//	})
	//}
	//return &V1.GetJobsResponse{
	//	Infos: infos,
	//}, err

}

func RunJob(c *gin.Context, req *V1.RunJobRequest) (*V1.RunJobResponse, error) {
	panic("implement me")
	//	var (
	//		err  error
	//		job  = &entity.Job{}
	//		tJob = &dao.TJob{}
	//		//infos []*params.GetJobsResponseInfo
	//	)
	//	if err != nil {
	//		goto ERR
	//	}
	//
	//	err = storage.GlobalMySQL.Table("t_jobs").Find(&tJob).Error
	//	if err != nil {
	//		return nil, err
	//	}
	//	job = &entity.Job{
	//		//ID:        tJob.ID,
	//		Targets:   strings.Split(tJob.ScriptIDs, ","),
	//		FilePath:  tJob.FilePath,
	//		Name:      tJob.Name,
	//		Type:      tJob.Type,
	//		ScriptIDS: strings.Split(tJob.ScriptIDs, ","),
	//	}
	//	go func() {
	//
	//		_ = job.Exec(c)
	//	}()
	//	return &V1.RunJobResponse{}, err
	//ERR:
	//	{
	//		return nil, err
	//	}
}
