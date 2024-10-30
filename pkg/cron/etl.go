package cron

import (
	"context"
	"goOrigin/pkg/logger"
	"time"
)

var log, _ = logger.InitZap()

// EsGather 是一个获取 Pod 信息的任务
type EsGather struct {
	name     string
	interval time.Time
}

// Exec 实现 Job 接口中的 Run 方法
func (p *EsGather) Exec(ctx context.Context) error {
	var (
		err    error
		ticker *time.Ticker = time.NewTicker(60 * time.Second)
		//esClient              = elastic.GlobalEsConns[config.ConfV2.Base.Mode]
	)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return err

		case <-ticker.C:
			// 获取 Pod 信息
		}
	}

}

// Name 实现 Job 接口中的 Name 方法
func (p *EsGather) Name() string {
	return p.name
}

func (p *EsGather) GetIanInfo(req byte) (string, error) {
	var (
	//query = map[string]interface{}{}
	//alias = "ian"
	)
	return "", nil
}

// 模拟获取 Pod 信息的函数

func EsGatherCronFactory() error {
	//var (
	//	err          error
	//	db           = mysql.GlobalMySQLConns[config.ConfV2.Base.Region]
	//	esCli        = elastic.GlobalEsConns[config.ConfV2.Base.Region]
	//	syncProjects []string
	//)
	//tRecords := db.Client.Table("projects").Select("project_id").Where("sync_status = ?", 1).Find(&syncProjects)
	//

	return nil
}
