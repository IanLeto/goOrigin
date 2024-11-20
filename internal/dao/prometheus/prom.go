package prometheus

import (
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"goOrigin/config"
)

var GlobalPrometheus = map[string]*PrometheusConn{}

type PrometheusConn struct {
	Client    *v1.API
	IsMigrate bool
}

func NewPromeV2Conns() *PrometheusConn {
	var (
		err error
	)
	conf := config.ConfV2
	for region, info := range conf.Env {
		GlobalPrometheus[region] = NewPrometheusConn(info.PromeConfig)
	}
	return &PrometheusConn{}

}

func NewPrometheusConn(conf config.PromeConfig) *PrometheusConn {
	var (
		err error
	)
	return &PrometheusConn{}

}
