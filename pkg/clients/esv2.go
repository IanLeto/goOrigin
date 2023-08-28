package clients

import (
	"bytes"
	"context"
	"encoding/json"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"goOrigin/config"
	"io"
)

var EsConns = map[string]*EsV2Conn{}

type EsV2Conn struct {
	Client *elasticsearch7.Client
}

func NewEsV2Conn(conf *config.EsInfo) *EsV2Conn {
	var (
		conn = &EsV2Conn{}
		err  error
	)

	client, err := elasticsearch7.NewClient(elasticsearch7.Config{
		Addresses: []string{
			conf.Address,
		},
	})
	if err != nil {
		panic(err)
	}
	conn.Client = client
	return conn
}

func (c *EsV2Conn) Creat(index string, body []byte) ([]byte, error) {
	var (
		//buf  bytes.Buffer
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

func (c *EsV2Conn) Query(index string, q map[string]interface{}) ([]byte, error) {
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

type EsDoc struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string      `json:"_index"`
			Type   string      `json:"_type"`
			Id     string      `json:"_id"`
			Score  float64     `json:"_score"`
			Source interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type InsertResultInfo struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	Id      string `json:"_id"`
	Version int    `json:"_version"`
	Result  string `json:"result"`
	Shards  struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	SeqNo       int `json:"_seq_no"`
	PrimaryTerm int `json:"_primary_term"`
}

func InitEs() error {
	for r, info := range config.Conf.Backend.EsConfig.ElasticSearchRegion {
		var ephemeral = info
		EsConns[r] = NewEsV2Conn(&ephemeral)
	}
	return nil
}
