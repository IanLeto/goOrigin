package cron

import (
	"context"
	"github.com/sirupsen/logrus"
	"goOrigin/pkg"
	"sync"
	"time"
)

// Job 接口，所有任务必须实现 Run 方法
type Job interface {
	Run() error
	Name() string
}

// GlobalCronTaskManager 管理任务的结构体
type GlobalCronTaskManager struct {
	jobChan       chan Job
	taskStatus    map[string]string
	statusMutex   sync.RWMutex
	tokenBucket   chan struct{}
	maxConcurrent int
	wg            sync.WaitGroup
	quit          chan struct{}
}

// 新建 GlobalCronTaskManager，初始化通道和令牌桶
func NewGlobalCronTaskManager(maxConcurrent int) *GlobalCronTaskManager {
	tm := &GlobalCronTaskManager{
		jobChan:       make(chan Job),
		taskStatus:    make(map[string]string),
		tokenBucket:   make(chan struct{}, maxConcurrent),
		maxConcurrent: maxConcurrent,
		quit:          make(chan struct{}),
	}
	// 初始化令牌桶
	for i := 0; i < maxConcurrent; i++ {
		tm.tokenBucket <- struct{}{}
	}
	return tm
}

// 启动任务管理器，监听 jobChan 并执行任务
func (tm *GlobalCronTaskManager) Start() {
	go func() {
		for {
			select {
			case job := <-tm.jobChan:
				tm.wg.Add(1)
				<-tm.tokenBucket // 获取令牌，控制并发

				// 更新任务状态为 "running"
				tm.setStatus(job.Name(), "running")

				// 使用 goroutine 执行任务
				go func(job Job) {
					defer tm.wg.Done()
					defer func() { tm.tokenBucket <- struct{}{} }() // 任务完成后归还令牌

					// 执行任务
					err := job.Run()
					if err != nil {
						tm.setStatus(job.Name(), "failed")
					} else {
						tm.setStatus(job.Name(), "completed")
					}
				}(job)

			case <-tm.quit:
				tm.wg.Wait() // 等待所有任务完成
				close(tm.jobChan)
				return
			}
		}
	}()
}

// 添加任务到任务管理器
func (tm *GlobalCronTaskManager) AddJob(job Job) {
	tm.setStatus(job.Name(), "waiting")
	tm.jobChan <- job
}

// 停止任务管理器
func (tm *GlobalCronTaskManager) Stop() {
	close(tm.quit)
}

// 设置任务状态
func (tm *GlobalCronTaskManager) setStatus(jobName string, status string) {
	tm.statusMutex.Lock()
	defer tm.statusMutex.Unlock()
	tm.taskStatus[jobName] = status
}

// 获取任务状态
func (tm *GlobalCronTaskManager) GetStatus(jobName string) string {
	tm.statusMutex.RLock()
	defer tm.statusMutex.RUnlock()
	if status, ok := tm.taskStatus[jobName]; ok {
		return status
	}
	return "not found"
}

// 获取所有任务的状态
func (tm *GlobalCronTaskManager) GetAllStatus() map[string]string {
	tm.statusMutex.RLock()
	defer tm.statusMutex.RUnlock()
	// 返回任务状态的副本
	statusCopy := make(map[string]string)
	for k, v := range tm.taskStatus {
		statusCopy[k] = v
	}
	return statusCopy
}

type CallBackFuncType func(t pkg.Job)

type TaskManager struct {
	// 待完成的任务队列, 外部可以不断往里推送任务
	JobQueue chan pkg.Job

	// worker并发任务的令牌桶
	tokenBucket chan interface{}
	maxWorker   int
	// 任务结束标志位
	ctx        context.Context
	cancelFunc context.CancelFunc
	// 回调函数的方法
	callBackFunc CallBackFuncType
	// 等待所有任务的结束的waitGroup
	wg sync.WaitGroup
	// 任务是否正在进行的标记位
	isRunning bool
	// 所有任务列表
	taskList  []pkg.Job
	jobWg     sync.WaitGroup
	startTime time.Time
	endTime   time.Time
}

func (m *TaskManager) PushTask() {
	// 开始执行任务
	m.startTime = time.Now()
	go func() {
		defer func() {
			// 任务关停
			m.isRunning = false
		}()
		for {
			select {
			case task := <-m.JobQueue:
				// 当我们拿到一个任务的时候，先去看看是否有可用的执行令牌
				<-m.tokenBucket
				// 拿到执行令牌，开始派遣goroutine 执行
				m.wg.Add(1)
				go func(job pkg.Job) {
					defer func() {
						// 本case 执行结束
						m.wg.Done()
						// 交还令牌
						m.tokenBucket <- struct{}{}
						// 总任务量
						m.jobWg.Done()
					}()
					if err := job.Exec(m.ctx, nil); err != nil {
						return
					}
					//m.callBackFunc(job.Run(m.ctx))
				}(task)
			case <-m.ctx.Done():
				return
			}
		}
	}()

	go func() {
		defer func() {
			m.jobWg.Done()
		}()
		for _, task := range m.taskList {
			select {
			case m.JobQueue <- task:
				m.jobWg.Add(1)
			// 考虑到有缓冲chan 关闭问题，这里选择使用ctx 主动关闭任务推送
			case <-m.ctx.Done():
				return
			}

		}
	}()
	//我们认为，push 任务到chan 也是需要时间的
	m.jobWg.Add(1)

}
func (m *TaskManager) PushTaskCallback() {
	// 开始执行任务
	m.startTime = time.Now()
	go func() {
		defer func() {
			// 任务关停
			m.isRunning = false
		}()
		for {
			select {
			case task := <-m.JobQueue:
				// 当我们拿到一个任务的时候，先去看看是否有可用的执行令牌
				<-m.tokenBucket
				// 拿到执行令牌，开始派遣goroutine 执行
				m.wg.Add(1)
				go func(job pkg.Job) {
					defer func() {
						// 本case 执行结束
						m.wg.Done()
						// 交还令牌
						m.tokenBucket <- struct{}{}
						// 总任务量
						m.jobWg.Done()
					}()
					m.callBackFunc(job)
				}(task)
			case <-m.ctx.Done():
				return
			}
		}
	}()

	go func() {
		defer func() {
			m.jobWg.Done()
		}()
		for _, task := range m.taskList {
			select {
			case m.JobQueue <- task:
				m.jobWg.Add(1)
			// 考虑到有缓冲chan 关闭问题，这里选择使用ctx 主动关闭任务推送
			case <-m.ctx.Done():
				return
			}
		}
	}()
	//我们认为，push 任务到chan 也是需要时间的
	m.jobWg.Add(1)

}

func (m *TaskManager) WaitJob() error {
	m.jobWg.Wait()
	// 关闭所有任务
	return nil
}
func (m *TaskManager) Wait() error {
	m.wg.Wait()
	// 关闭所有任务
	return nil
}

func (m *TaskManager) Stop() error {
	m.cancelFunc()
	logrus.Debugf("耗时%ds", time.Since(m.startTime))
	return nil
}

// NewTaskManager ctx , cancelctx, 任务列表, 任务队列， 令牌桶， 回调函数，如何执行
func NewTaskManager(ctx context.Context, cancelFunc context.CancelFunc, taskList []pkg.Job,
	jobQueue chan pkg.Job, tokenBucket chan interface{}, callBack CallBackFuncType) *TaskManager {
	return &TaskManager{
		ctx:          ctx,
		cancelFunc:   cancelFunc,
		taskList:     taskList,
		JobQueue:     jobQueue,
		tokenBucket:  tokenBucket,
		callBackFunc: callBack,
		jobWg:        sync.WaitGroup{},
		wg:           sync.WaitGroup{},
	}
}
