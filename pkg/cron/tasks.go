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
	RedisClient goRedis.Client
	Pipeline    goRedis.Pipeline
}

func NewCacheTask(interval time.Time) *RedisTask {
	cli := goRedis.NewClient(&goRedis.Options{
		Network:            "tcp",
		Addr:               "",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
	})
	return &RedisTask{
		RedisClient: goRedis.Client{},
		Pipeline:    goRedis.Pipeline{},
	}
}

// demo 定时redis 任务
type DemoTask struct {
	RedisClient goRedis.Client
	Pipeline    goRedis.Pipeline
}

func (d *DemoTask) Run(ctx context.Context) error {
	panic("implement me")
}

func NewDemoTask() (*DemoTask, error) {
	return &DemoTask{
		RedisClient: goRedis.Client{},
		Pipeline:    goRedis.Pipeline{},
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
