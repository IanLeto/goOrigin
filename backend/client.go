package backend

type Client interface {
	Close() error
	ShowType() string
}

type BaseClient struct {
	ClientType string
}

func (b BaseClient) Close() error {
	panic("implement me")
}

func (b BaseClient) NewClient() interface{} {
	panic("implement me")
}

func (b BaseClient) ShowType() string {
	return b.ClientType
}
