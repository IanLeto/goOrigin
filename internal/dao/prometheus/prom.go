package prometheus

import (
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"goOrigin/config"
	"goOrigin/pkg/logger"
)

var GlobalPrometheus = map[string]*PrometheusConn{}
var log, _ = logger.InitZap()

type PrometheusConn struct {
	Client    v1.API
	IsMigrate bool
}

func NewPromeV2Conns() error {
	var (
		err error
	)
	conf := config.ConfV2
	for region, info := range conf.Env {
		GlobalPrometheus[region], err = NewPrometheusConn(info.PromeConfig)
		if err != nil {
			log.Sugar().Errorf("prometheus conn error: %v", err)
		}
	}
	return err
}

func NewPrometheusConn(conf config.PromeConfig) (*PrometheusConn, error) {
	var (
		err error
	)
	client, err := api.NewClient(api.Config{
		Address: conf.Address,
	})

	return &PrometheusConn{
		Client: v1.NewAPI(client),
	}, err
}
