package moniter

import "github.com/prometheus/client_golang/prometheus"

var SuccessCount = newSuccessCountVec()

func newSuccessCountVec() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "success span count",
		Help: "成功的span count",
	}, []string{"trans_type_code", "trans_channel", "ret_code", "svc_name"})
}

func newFailedCountGaugeVec() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "Failed span count",
		Help: "失败的span count",
	}, []string{"trans_type_code", "trans_channel", "ret_code", "svc_name"})
}
