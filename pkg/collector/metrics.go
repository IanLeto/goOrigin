package collector

import "github.com/prometheus/client_golang/prometheus"

var APICounterVec *prometheus.CounterVec

func init() {
	APICounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "ian_test_api",
		Help: "hei",
	}, []string{"api"})
	prometheus.MustRegister(APICounterVec)
	
}
