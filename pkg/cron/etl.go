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

// NewSyncDataGlobalJob
// 1. 初始化任务
// 2. 本质是个独立的goroutinue，和ticker 添加到任务队列
func NewSyncDataGlobalJob() error {

	go func() {
		var (
			dbCli = mysql.GlobalMySQLConns[config.ConfV2.Base.Region]
			esCli = elastic.GlobalEsConns[config.ConfV2.Base.Region]
			interval,err = time.ParseDuration(config.ConfV2.Cron["SyncDataJob"].)
		)
		for  {
			select {
			case <-time.NewTimer(60 * time.Second).C:
			default:
				return

			}
		}
		GlobalSyncDataJob = &SyncDataJob{
			name:     "SyncDataJob",
			interval: config.ConfV2.Cron["SyncDataJob"].(time.Time),
			dbCli:    *dbCli,
			esCli:    *esCli,
		}
		GTM.AddJob(GlobalSyncDataJob)
	}()

	return nil
}
