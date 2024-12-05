package cron

import (
	"context"
	"fmt"
	"goOrigin/pkg"
	//"goOrigin/pkg/logger"
	"time"
)

type PodFactory struct {
	Trick   *time.Ticker
	Retry   int
	Timeout time.Duration
	Manger  *TaskManager
}

// Exec 将任务添加到任务队列，任务队列全局唯一
func (l *PodFactory) Exec(rootCtx context.Context, info pkg.JobMessageInfo) error {
	var (
		logger2 = logger
		err     error
	)
	if err != nil {
		logger2.Error("日志生成的cron 失败")
		return nil
	}
	var (
		jobQueue    = make(chan pkg.Job, 50)
		tokenBucket = make(chan interface{}, 50)
		jobList     []pkg.Job
		ctx, cancel = context.WithCancel(rootCtx)
	)
	// 填充令牌桶
	for i := 0; i < 50; i++ {
		tokenBucket <- struct{}{}
	}

	l.Manger = NewTaskManager(ctx, cancel, jobList, jobQueue, tokenBucket, func(job pkg.Job) {
		_ = job.Exec(ctx, nil)
	})

	for {
		select {
		case <-l.Trick.C:
			for i := 0; i < 1; i++ {
				jobList = append(jobList, &PodInfoJob{})
			}
			l.Manger.PushTask()
			//time.Sleep(10 * time.Second)
			//// 到时间后，将任务添加到任务队列
			//l.Manger.JobQueue <- &PodInfoJob{}
			// 推送任务
		}
	}
}

func (l *PodFactory) Stop(ctx context.Context, kill chan struct{}) error {
	//TODO implement me
	panic("implement me")
}
func RegPodInfoCronFactory() error {
	task := &PodFactory{
		Trick: time.NewTicker(10 * time.Second),
	}
	QueueCron = append(QueueCron, task)
	return nil

}

func NewPodTaskPool() (*TaskManager, error) {
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
		jobList = append(jobList, &PodFactory{})
	}
	manager := NewTaskManager(ctx, cancel, jobList, jobQueue, tokenBucket, func(job pkg.Job) {
		_ = job.Exec(ctx, nil)
	})

	return manager, nil
}

type PodInfoJob struct {
}

func (p *PodInfoJob) Exec(ctx context.Context, info pkg.JobMessageInfo) error {
	fmt.Println("todo")
	return nil
}

func (p *PodInfoJob) Stop(ctx context.Context, kill chan struct{}) error {
	//TODO implement me
	panic("implement me")
}
