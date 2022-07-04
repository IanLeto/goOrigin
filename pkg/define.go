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

type JobMessageInfo interface {
	Export()
}

// Job 任务执行的最小单元
type Job interface {
	Exec(ctx context.Context, info JobMessageInfo) error
	Stop(ctx context.Context, kill chan struct{}) error
}
