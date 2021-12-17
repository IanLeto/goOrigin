package pkg

import (
	"context"
)

type Conn interface {
	Close() error
	Exec() ([]byte, error)
}

type IClient interface {
	Close() error
	Ping() error
}

// 任务大类
type Task interface {
	Run(ctx context.Context) error
}

type JobMessageInfo interface {
	Export()
}

// 任务执行的最小单元
type Job interface {
	Exec(ctx context.Context, info JobMessageInfo) error
	Stop(ctx context.Context, kill chan struct{}) error
}
