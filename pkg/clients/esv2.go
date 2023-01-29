package clients

import (
	"bytes"
	"encoding/json"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"goOrigin/config"
	"io"
)

type EsV2Conn struct {
	Client *elasticsearch7.Client
}

func NewEsV2Conn(conf *config.Config) *EsV2Conn {
	var (
		conn = &EsV2Conn{}
		err  error
	)
	if conf == nil {
		conf = config.Conf
	}
	client, err := elasticsearch7.NewClient(elasticsearch7.Config{
		Addresses: []string{
			conf.Backend.EsConfig.Address,
		},
	})
	if err != nil {
		panic(err)
	}
	conn.Client = client
	return conn
}

func (c *EsV2Conn) Query(index string, q map[string]interface{}) (*Response, error) {
	res, err := c.Client.Info()
	if err != nil {
		return nil, err
	}
	var (
		buf        bytes.Buffer
		resp       Response
		resContent []byte
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
	resContent, err = io.ReadAll(res.Body)
	err = json.Unmarshal(resContent, &resp)
	return &resp, err

ERR:
	{
		return nil, err
	}
}

type Response struct {
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
