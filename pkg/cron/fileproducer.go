package cron

import (
	"context"
	"fmt"
	"goOrigin/pkg"
	"goOrigin/pkg/logger"
	"math/rand"
	"time"
)

type LoggerProducer struct {
	Trick *time.Ticker
}

func (l *LoggerProducer) Exec(ctx context.Context, info pkg.JobMessageInfo) error {
	var (
		logger2 = logger.NewLogger()
		err     error
	)
	if err != nil {
		logger2.Error("日志生成的cron 失败")
		return nil
	}
	for {
		select {
		case <-l.Trick.C:
			randNum := rand.Intn(10)
			switch randNum {
			// 随机生成 error and warn and info
			case 8:
				logger2.Error("fake error log")
			case 5:
				logger2.Warn("fake warn log")
			default:
				logger2.Info(fmt.Sprintf("log test %s", time.Now().String()))
			}

		}
	}
}

func (l *LoggerProducer) Stop(ctx context.Context, kill chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func RegLoggerCron() error {
	task := &LoggerProducer{
		time.NewTicker(10 * time.Second),
	}
	QueueCron = append(QueueCron, task)
	return nil

}
