package clients

type Conn interface {
	Migrate() error
}
