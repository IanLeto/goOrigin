package clients

import (
	"fmt"
	"github.com/prometheus/client_golang/api"
	"goOrigin/config"
	"goOrigin/pkg/logger"
)

func NewPromClient() (api.Client, error) {
	var (
		err error
		log = logger.Logger
	)
	client, err := api.NewClient(api.Config{
		Address: config.Conf.Backend.PromConfig.Address,
	})
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		return nil, err
	}
	return client, err
}
