package cron

import (
	"context"
	"fmt"
	"goOrigin/config"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"time"
)

var log = logger
var GlobalSyncDataJob *SyncDataJob

// SyncDataJob 是一个获取 Pod 信息的任务
type SyncDataJob struct {
	name     string
	interval time.Time
	dbCli    mysql.MySQLConn
	esCli    elastic.EsV2Conn
	project  string
}

// Exec 实现 Job 接口中的 Run 方法
func (p *SyncDataJob) Exec(ctx context.Context) error {
	//var (
	//	err    error
	//	ticker *time.Ticker = time.NewTicker(60 * time.Second)
	//)
	fmt.Println("执行中", time.Now())
	time.Sleep(1 * time.Second)
	return nil

}

// Title 实现 Job 接口中的 Title 方法
func (p *SyncDataJob) GetName() string {
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
// 2. 本质是个独立的goroutinue，和ticker 添加到任务队列,每隔一段时间，推送任务到任务队列
func NewSyncDataGlobalJob() error {
	log.Info("NewSyncDataGlobalJob 启动")
	var (
		interval int
		dbCli    *mysql.MySQLConn
		esCli    *elastic.EsV2Conn
		projects []string
	)
	interval = config.ConfV2.Env[config.ConfV2.Base.Region].CronJobConfig.TransferConfig.Interval
	dbCli = mysql.GlobalMySQLConns[config.ConfV2.Base.Region]
	esCli = elastic.GlobalEsConns[config.ConfV2.Base.Region]

	// todo
	//dbCli.Client.Select(&projects, "select project from project")
	interval = 10

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic in goroutine:", r)
			}
		}()
		for {
			select {
			case <-time.NewTimer(time.Duration(interval) * time.Second).C:
				// 获取 Pod 信息
				for _, p := range projects {
					GTM.AddJob(&SyncDataJob{
						name:     fmt.Sprintf("SyncDataJob-%s", config.ConfV2.Base.Region),
						interval: time.Now(),
						dbCli:    *dbCli,
						esCli:    *esCli,
						project:  p,
					})
				}

			}
		}

	}()

	return nil
}

func RegTransfer(ctx context.Context) error {
	// 创建一个 goroutine 来注册和管理任务
	go func() {
		// 使用 defer 捕获 goroutine 中的 panic
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic in RegTransfer goroutine:", r)
			}
		}()

		// 启动 NewSyncDataGlobalJob
		err := NewSyncDataGlobalJob()
		if err != nil {
			fmt.Println("Error starting NewSyncDataGlobalJob:", err)
			return
		}

		// 监听 context 的取消信号
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled, stopping RegTransfer task")
			// 当 context 被取消时，退出此 goroutine
			return
		}
	}()

	return nil
}
