package cron_test

import (
	"context"
	"fmt"
	"goOrigin/config"
	"goOrigin/pkg"
	"goOrigin/pkg/cron"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestDemoJob struct {
}

func NewTestDemoJob() *TestDemoJob {
	return &TestDemoJob{}
}

func NewDemoTaskPool() (*cron.TaskManager, error) {
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
		jobList = append(jobList, NewTestDemoJob())
	}
	manager := cron.NewTaskManager(ctx, cancel, jobList, jobQueue, tokenBucket, func(job pkg.Job) {
		_ = job.Exec(ctx, nil)
	})

	return manager, nil
}

func (t *TestDemoJob) Exec(ctx context.Context, info pkg.JobMessageInfo) error {
	fmt.Println(rand.Intn(100))
	return nil
}

func (t *TestDemoJob) Stop(ctx context.Context, kill chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

// RedisSuite :
type JobSuite struct {
	suite.Suite
	conf *config.Config
	jobs *cron.TaskManager
}

func (s *JobSuite) SetupTest() {
	s.jobs, _ = NewDemoTaskPool()
}

// TestMarshal :
func (s *JobSuite) TestConfig() {
	s.jobs.PushTask()
	s.NoError(s.jobs.Wait())
	s.NoError(s.jobs.WaitJob())

}

// TestHttpClient :
func TestRedisConfiguration(t *testing.T) {
	suite.Run(t, new(JobSuite))
}
