package cron

import (
	"context"
	"fmt"
	goRedis "github.com/go-redis/redis"
	"goOrigin/config"
	"goOrigin/pkg"
	"math/rand"
	"time"
)

var QueueCron []pkg.Job

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

// demo 定时redis 任务 定时更新某个key
type DemoTask struct {
	Manager *TaskManager
}
type DemoJob struct {
}

func NewDemoJob() *DemoJob {
	return &DemoJob{}
}

func (d *DemoJob) Exec(ctx context.Context, info pkg.JobMessageInfo) error {
	time.Sleep(1 * time.Second)
	fmt.Println(rand.Intn(50))
	return nil
}

func (d *DemoJob) Stop(ctx context.Context, kill chan struct{}) error {
	panic("implement me")
}

func RegisterDemoTask() error {
	DemoJob := &DemoJob{}
	QueueCron = append(QueueCron, DemoJob)
	return nil
}

func NewDemoTask() (*DemoTask, error) {
	var (
		jobQueue    = make(chan pkg.Job, 50)
		tokenBucket = make(chan interface{}, 50)
		jobList     []pkg.Job
		ctx, cancel = context.WithCancel(context.Background())
	)
	// 填充令牌桶
	for i := 0; i < 50; i++ {
		tokenBucket <- struct{}{}
	}
	for i := 0; i < 1000; i++ {
		jobList = append(jobList, NewDemoJob())
	}
	manager := NewTaskManager(ctx, cancel, jobList, jobQueue, tokenBucket, func(job pkg.Job) {
		_ = job.Exec(ctx, nil)
	})

	return &DemoTask{
		Manager: manager,
	}, nil
}

func (d *DemoTask) Run(ctx context.Context) error {
	d.Manager.PushTaskCallback()

	if err := d.Manager.Wait(); err != nil {
		return err
	}
	if err := d.Manager.WaitJob(); err != nil {
		return err
	}
	return nil
}

// ______________________________________________________________

// 心跳
type HearBeatTask struct {
}

func (h *HearBeatTask) Run(ctx context.Context) error {
	panic("implement me")
}

func NewHearBeatTask() (*HearBeatTask, error) {
	return &HearBeatTask{}, nil
}

func RegisterHearBeatTask() error {

	QueueCron = append(QueueCron, nil)
	return nil
}
