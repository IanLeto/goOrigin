package clients

import (
	"github.com/olivere/elastic/v7"
	"goOrigin/config"
)

type EsAPIs interface {
	Ping() error
}

type EsConn struct {
	Client *elastic.Client
}

func NewEsConn(conf *config.Config) (*EsConn, error) {
	var (
		conn = &EsConn{}
		err  error
	)
	if conf == nil {
		conf = config.Conf
	}
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetHealthcheck(false),
		elastic.SetURL(config.Conf.Backend.EsConfig.Address))
	if err != nil {
		return nil, err
	}
	conn.Client = client
	return conn, err
}

func NewESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetHealthcheck(false),
		elastic.SetURL(config.Conf.Backend.EsConfig.Address))
	if err != nil {
		return nil, err
	}
	return client, err
}

func (c *EsConn) Ping() error {
	return nil
}

func (c *EsConn) Search(callback func()) (interface{}, error) {
	panic(1)
}

func (c *EsConn) Create(callback func()) (interface{}, error) {
	panic(1)
}


func (c *EsConn) Update(callback func()) (interface{}, error) {
	panic(1)
}

func (c *EsConn) Delete(callback func()) (interface{}, error) {
	panic(1)
}



