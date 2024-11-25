package moniter

import "github.com/prometheus/client_golang/prometheus"

var IanRecordWeightGaugeVec = newIanRecordWeightGaugeVec()

func newIanRecordWeightGaugeVec() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "weight",
		Help: "ian weight gau per day",
		ConstLabels: map[string]string{
			"region": "",
			"time":   "",
		},
	}, []string{})
}

func SetWeightGauge(metric float64, region string, time string) {
	IanRecordWeightGaugeVec.With(prometheus.Labels{
		"region": region,
		"time":   time,
	}).Set(metric)
}
