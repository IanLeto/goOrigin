package moniter

import "github.com/prometheus/client_golang/prometheus"

func RegMoniter(cols ...prometheus.Collector) {

	prometheus.MustRegister(cols...)
}
func init() {
	cpu := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ian_test_cpu",
		Help: "hei",
	})
	cpu.Set(1)
	RegMoniter(cpu)
}
