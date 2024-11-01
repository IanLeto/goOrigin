package cron

import (
	"context"
	"goOrigin/config"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"goOrigin/pkg/logger"
	"time"
)

var log, _ = logger.InitZap()
var GlobalSyncDataJob *SyncDataJob

// SyncDataJob 是一个获取 Pod 信息的任务
type SyncDataJob struct {
	name     string
	interval time.Time
	dbCli    mysql.MySQLConn
	esCli    elastic.EsV2Conn
}

// Exec 实现 Job 接口中的 Run 方法
func (p *SyncDataJob) Exec(ctx context.Context) error {
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
func (p *SyncDataJob) Name() string {
	return p.name
}

func (p *SyncDataJob) GetIanInfo(req byte) (string, error) {
	var (
	//query = map[string]interface{}{}
	//alias = "ian"
	)
	return "", nil
}

// NewSyncDataGlobalJob  模拟获取 Pod 信息的函数；初始化任务本身
func NewSyncDataGlobalJob() error {
	var (
		dbCli = mysql.GlobalMySQLConns[config.ConfV2.Base.Region]
		esCli = elastic.GlobalEsConns[config.ConfV2.Base.Region]
	)

	GlobalSyncDataJob = &SyncDataJob{
		name:     "SyncDataJob",
		interval: config.ConfV2.Cron["SyncDataJob"].(time.Time),
		dbCli:    *dbCli,
		esCli:    *esCli,
	}
	GTM.AddJob(GlobalSyncDataJob)
	return nil
}
