package pkg

type Conn interface {
	Close() error
	Exec() ([]byte, error)
}
