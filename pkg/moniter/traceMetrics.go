package moniter

import (
	"github.com/prometheus/client_golang/prometheus"
	"goOrigin/internal/model/entity"
	// 假设你的结构体定义在 entity 包中
)

// 定义 Prometheus 指标
var (
	SuccessCountGaugeVec  = newSuccessCountGaugeVec()
	FailedCountGaugeVec   = newFailedCountGaugeVec()
	ResponseCountGaugeVec = newResponseCountGaugeVec()
	SuccessRateGaugeVec   = newSuccessRateGaugeVec()
	FailedRateGaugeVec    = newFailedRateGaugeVec()
)

// 初始化 SuccessCount 指标
func newSuccessCountGaugeVec() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "success_count",
		Help: "Number of successful transactions",
	}, []string{"cluster", "trans_type", "trans_type_code", "trans_channel", "svc_name", "project_name", "pod_name"})
}

// 初始化 FailedCount 指标
func newFailedCountGaugeVec() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "failed_count",
		Help: "Number of failed transactions",
	}, []string{"cluster", "trans_type", "trans_type_code", "trans_channel", "svc_name", "project_name", "pod_name"})
}

// 初始化 ResponseCount 指标
func newResponseCountGaugeVec() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "response_count",
		Help: "Total number of responses",
	}, []string{"cluster", "trans_type", "trans_type_code", "trans_channel", "svc_name", "project_name", "pod_name"})
}

// 初始化 SuccessRate 指标
func newSuccessRateGaugeVec() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "success_rate",
		Help: "Rate of successful transactions",
	}, []string{"cluster", "trans_type", "trans_type_code", "trans_channel", "svc_name", "project_name", "pod_name"})
}

// 初始化 FailedRate 指标
func newFailedRateGaugeVec() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "failed_rate",
		Help: "Rate of failed transactions",
	}, []string{"cluster", "trans_type", "trans_type_code", "trans_channel", "svc_name", "project_name", "pod_name"})
}

// 注册 Prometheus 指标并初始化部分默认值
func init() {
	Reg.MustRegister(SuccessCountGaugeVec)
	Reg.MustRegister(FailedCountGaugeVec)
	Reg.MustRegister(ResponseCountGaugeVec)
	Reg.MustRegister(SuccessRateGaugeVec)
	Reg.MustRegister(FailedRateGaugeVec)

	// 初始化默认值
	initializeMetrics()
}

// 初始化指标的默认值
func initializeMetrics() {
	defaultLabels := prometheus.Labels{
		"cluster":         "global",
		"trans_type":      "10ff1",
		"trans_type_code": "0000",
		"trans_channel":   "FF01",
		"svc_name":        "cpaas",
		"project_name":    "cpaas",
		"pod_name":        "iantest",
	}

	// 为每个指标设置默认值
	SuccessCountGaugeVec.With(defaultLabels).Set(10)   // 成功次数默认值
	FailedCountGaugeVec.With(defaultLabels).Set(2)     // 失败次数默认值
	ResponseCountGaugeVec.With(defaultLabels).Set(12)  // 响应总数默认值
	SuccessRateGaugeVec.With(defaultLabels).Set(83.33) // 成功率默认值 (10/12*100)
	FailedRateGaugeVec.With(defaultLabels).Set(16.67)  // 失败率默认值 (2/12*100)
}

// **从结构体设置指标**
func SetODAMetrics(metric *entity.ODAMetricEntity) {
	labels := prometheus.Labels{
		"cluster":         metric.Cluster,
		"trans_type_code": metric.TransTypeCode,
		"trans_channel":   metric.TransChannel,
		"svc_name":        metric.SvcName,
	}

	// 设置指标值
	SuccessCountGaugeVec.With(labels).Set(float64(metric.SuccessCount))
	FailedCountGaugeVec.With(labels).Set(float64(metric.FailedCount))
	ResponseCountGaugeVec.With(labels).Set(float64(metric.ResponseCount))
	SuccessRateGaugeVec.With(labels).Set(float64(metric.SuccessRate))
	FailedRateGaugeVec.With(labels).Set(float64(metric.FailedRate))
}
