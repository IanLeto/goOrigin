package moniter

import "github.com/prometheus/client_golang/prometheus"

var IanRecordWeightGaugeVec = newIanRecordWeightGaugeVec()

func newIanRecordWeightGaugeVec() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "weight",
		Help: "ian weight gau per day",
	}, []string{"region", "time"})
}

var SpanCount = newSpanCount()

func newSpanCount() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "success span count",
		Help: "",
	}, []string{"region", "time"})
}
func SetWeightGauge(metric float64, region string, time string) {
	IanRecordWeightGaugeVec.With(prometheus.Labels{
		"region": region,
		"time":   time,
	}).Set(metric)
}

func init() {
	Reg.MustRegister(IanRecordWeightGaugeVec)
	Reg.MustRegister(SpanCount)
	SetWeightGauge(100, "sa", "xs")
}
