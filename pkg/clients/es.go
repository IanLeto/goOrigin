package clients

import (
	"github.com/olivere/elastic/v7"
)

type EsAPIs interface {
	Ping() error
}

type EsConn struct {
	Client *elastic.Client
}

func NewESClient() (*elastic.Client, error) {
	panic(1)
}

func (c *EsConn) Ping() error {
	return nil
}

func (c *EsConn) Search(callback func()) (interface{}, error) {
	//res, err := c.Client.

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
