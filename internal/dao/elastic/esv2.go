package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
)

var (
	LogIndict = ``
	EventDict = ``
)

var GlobalEsConns = map[string]*EsV2Conn{}

type EsV2Conn struct {
	Client *elasticsearch7.Client
}

func (c *EsV2Conn) Migrate() error {
	//c.Client.
	panic(1)
}

func (c *EsV2Conn) Create(index string, body []byte) ([]byte, error) {
	var (
		req  = esapi.IndexRequest{}
		resp *esapi.Response
		err  error
	)

	req = esapi.IndexRequest{
		Index: index,
		Body:  bytes.NewReader(body),
	}
	resp, err = req.Do(context.TODO(), c.Client)
	if err != nil {
		goto ERR
	}
	defer func() { _ = resp.Body.Close() }()

	return io.ReadAll(resp.Body)

ERR:
	{
		return nil, err
	}

}

func (c *EsV2Conn) Search(index string, q map[string]interface{}) ([]byte, error) {
	res, err := c.Client.Info()
	if err != nil {
		return nil, err
	}
	var (
		buf bytes.Buffer
	)

	err = json.NewEncoder(&buf).Encode(q)
	if err != nil {
		goto ERR
	}
	res, err = c.Client.Search(
		c.Client.Search.WithIndex(index),
		c.Client.Search.WithBody(&buf),
	)
	if err != nil {
		goto ERR
	}
	defer func() { _ = res.Body.Close() }()

	return io.ReadAll(res.Body)

ERR:
	{
		return nil, err
	}
}
