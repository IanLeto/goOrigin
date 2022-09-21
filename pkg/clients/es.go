package clients

import (
	"github.com/olivere/elastic/v7"
	"goOrigin/config"
)

func NewESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetHealthcheck(false),
		elastic.SetURL(config.Conf.Backend.EsConfig.Address))
	if err != nil {
		return nil, err
	}
	return client, err
}
