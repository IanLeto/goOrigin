package cron

import (
	"context"
	goRedis "github.com/go-redis/redis"
	"goOrigin/config"
	"goOrigin/pkg"
	"time"
)

var QueueCron []pkg.Task

// 定时刷新缓存
type CacheTask interface {
}

// redis cache
type RedisTask struct {
	RedisClient *goRedis.Client
	Pipeline    goRedis.Pipeliner
}

func NewCacheTask(interval time.Time) *RedisTask {
	cli := goRedis.NewClient(&goRedis.Options{
		Network: "tcp",
		Addr:    config.Conf.Backend.RedisConfig.Addr,
	})
	pipe := cli.Pipeline()
	return &RedisTask{
		RedisClient: cli,
		Pipeline:    pipe,
	}
}

// demo 定时redis 任务
type DemoTask struct {
	RedisClient *goRedis.Client
	Pipeline    goRedis.Pipeliner
}

func (d *DemoTask) Run(ctx context.Context) error {
	panic("implement me")
}

func NewDemoTask() (*DemoTask, error) {
	cli := goRedis.NewClient(&goRedis.Options{
		Network: "tcp",
		Addr:    config.Conf.Backend.RedisConfig.Addr,
	})
	pipe := cli.Pipeline()
	return &DemoTask{
		RedisClient: cli,
		Pipeline:    pipe,
	}, nil
}

func RegisterDemoTask() error {
	demoTask, _ := NewDemoTask()
	QueueCron = append(QueueCron, demoTask)
	return nil
}

// 心跳
type HearBeatTask struct {
}

func (h *HearBeatTask) Run(ctx context.Context) error {
	panic("implement me")
}

func NewHearBeatTask() (*HearBeatTask, error) {
	return &HearBeatTask{
	}, nil
}

func RegisterHearBeatTask() error {
	demoTask, _ := NewDemoTask()
	QueueCron = append(QueueCron, demoTask)
	return nil
}
