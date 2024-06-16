package entity

import (
	"goOrigin/internal/model/dao"
	"strings"
)

type Job struct {
	ID         uint
	Targets    []string
	FilePath   string
	Name       string
	Type       string
	StrategyID uint
	Scripts    BaseScript
	ScriptIDS  []string
}

func (j *Job) ToTable() *dao.TJob {
	return &dao.TJob{
		Name:      j.Name,
		Target:    strings.Join(j.Targets, ","),
		FilePath:  j.FilePath,
		Type:      j.Type,
		ScriptIDs: strings.Join(j.ScriptIDS, ","),
	}
}

//func (j *Job) Create() error {
//	tJob := j.ToTable()
//	_, err := tJob.Create()
//	if err != nil {
//		return err
//	}
//	//j.ID = tJob.ID
//	return nil
//}
//
//func (j *Job) Update() error {
//	tJob := j.ToTable()
//	_, err := tJob.Update()
//	if err != nil {
//		return err
//	}
//	//j.ID = tJob.ID
//	return nil
//}
//
//func (j *Job) Delete() error {
//	return dao.DeleteJobByID(j.ID)
//}
//
//func (j *Job) QueryDetail() (*dao.TJob, error) {
//	tJob := j.ToTable()
//	err := storage.GlobalMySQL.Model(tJob).First(tJob).Error
//	if err != nil {
//		return nil, err
//	}
//	return tJob, err
//}
//
//func (j *Job) Exec(ctx context.Context) error {
//	return nil
//	//	var (
//	//		err     error
//	//		log     = logger.NewLogger()
//	//		client  *elastic.Client
//	//		scripts []*BaseScript
//	//	)
//	//	agent, err := rpcClient.NewAgentClient()
//	//	if err != nil {
//	//		log.Error(fmt.Sprintf("agent 创建失败 %s", err))
//	//		goto ERR
//	//	}
//	//	if j.FilePath != "" {
//	//		return err
//	//	}
//	//	client, err = clients.NewESClient()
//	//	if err != nil {
//	//		log.Error(fmt.Sprintf("es 创建失败 %s", err))
//	//		goto ERR
//	//	}
//	//	for _, id := range j.ScriptIDS {
//	//		// 先全部遍历出来
//	//		res, err := BoolQueryScript(ctx, client, elastic.Search(elastic.NewBoolQuery().Filter(elastic.NewTermQuery("ID", id))))
//	//		if err != nil {
//	//			log.Error(fmt.Sprintf("查询es client 失败 %s", err.Error()))
//	//			goto ERR
//	//		}
//	//		scripts = append(scripts, res...)
//	//	}
//	//	for _, i := range scripts {
//	//		_, err = agent.RunScript(ctx, &pbs.RunScriptRequest{
//	//			Name:    i.Name,
//	//			Content: i.Content,
//	//		})
//	//		if err != nil {
//	//			log.Error(fmt.Sprintf("脚本%s 执行失败 %s", i.Name, err))
//	//			goto ERR
//	//		}
//	//	}
//	//	return err
//	//ERR:
//	//	{
//	//		return err
//	//	}
//}
