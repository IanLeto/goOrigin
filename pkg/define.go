package pkg

type Conn interface {
	Close() error
	Exec() ([]byte, error)
}

type IClient interface {
	Close() error
	Ping() error

}
