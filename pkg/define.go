package pkg

import "context"

type Conn interface {
	Close() error
	Exec() ([]byte, error)
}

type IClient interface {
	Close() error
	Ping() error
}

type Task interface {
	Run(ctx context.Context) error
}
