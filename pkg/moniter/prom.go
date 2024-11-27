package moniter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Reg = prometheus.NewRegistry()

func init() {
	cpu := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ian_test_cpu",
		Help: "hei",
	})
	cpu.Set(1)
	Reg.MustRegister(cpu)

}
