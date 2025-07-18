package entity

import (
	"fmt"
	"time"
)

type ODAMetricEntity struct {
	Interval time.Duration `json:"interval"`
	*PredefinedDimensions
	*Indicator
	CustomDimensions []string `json:"custom_dimensions"`
}

type PredefinedDimensions struct {
	TraceID       string `json:"trace_id"`
	Cluster       string `json:"cluster"`
	TransTypeCode string `json:"trans_type_code"` // 锚定字段
	TransChannel  string `json:"trans_channel"`
	RetCode       string `json:"ret_code"`
	SvcName       string `json:"svc_name"`
}

type Indicator struct {
	SuccessCount  int `json:"success_count"`
	SuccessRate   int `json:"success_rate"`
	FailedCount   int `json:"failed_count"`
	FailedRate    int `json:"failed_rate"`
	ResponseCount int `json:"response_count"`
	ResponseRate  int `json:"response_rate"`
}

type SpanEntity struct {
	Stats []SpanDatEntity `json:"stats"`
}

type SpanDatEntity struct {
	TransType    string `json:"trans_type"`
	TransTypeCN  string `json:"trans_type_cn"`
	SuccessCount int64  `json:"success_count"`
	FailedCount  int64  `json:"failed_count"`
	UnknownCount int64  `json:"unknown_count"`
	Total        int64  `json:"total"`
}

func ConvertLogToMetric(log *KafkaLogEntity) ODAMetricEntity {
	// 组装目标结构
	metric := ODAMetricEntity{
		PredefinedDimensions: &PredefinedDimensions{
			Cluster:       log.InstanceZone,
			TransTypeCode: log.LogType,
			TransChannel:  log.RemoteApp,
			RetCode:       log.ResultCode,
			SvcName:       log.Service,
		},
	}
	return metric
}

// TransInfoEntity 修改后的实体结构体
type TransInfoEntity struct {
	Project     string              `json:"project"`
	TransType   string              `json:"trans_type"`    // 等价于 url_path
	TransTypeCn []string            `json:"trans_type_cn"` // 改为数组，存储多个中文名称
	ReturnCodes []*ReturnCodeEntity `json:"return_codes"`
	IsAlert     bool                `json:"is_alert"`
	Threshold   int                 `json:"threshold"`
}
type ReturnCodeEntity struct {
	ReturnCode string `json:"return_code"`
	Project    string `json:"project"`
	TransType  string `json:"trans_type"`
	Status     string `json:"status"`
	Count      int    `json:"count"` // 存储计数
}

type TradeReturnCodeEntity struct {
	UrlPath       string
	SuccessCount  int
	FailedCount   int
	UnKnownCount  string
	Total         string
	TransTypeCn   string
	ResponseCount string
}

type TransTypeEntity struct {
	TransType   string   `json:"trans_type"`
	TransTypeCn string   `json:"trans_type_cn"`
	ReturnCodes []string `json:"return_codes"`
}

type TransTypeResponseEntity struct {
	Items []*TransTypeEntity `json:"items"`
}

type UrlPathAggEntity struct {
	TransType       string            `json:"trans_type"`        // 等价于 trans_type
	TransTypeCN     string            `json:"trans_type_cn"`     // 等价于 trans_type_cn
	Project         string            `json:"project"`           // 新增：与TransInfoEntity对应
	ReturnCode      map[string]string `json:"return_code"`       // key: return_code, value: return_code_cn
	ReturnCodeCount map[string]int    `json:"return_code_count"` // key: return_code, value: count
	Interval        int               `json:"interval"`          // 新增：与TransInfoEntity对应
}

type TradePubMessageEntity struct {
}

// TransTypeKey 用于查询的键值对
type TransTypeKey struct {
	TransType string `json:"trans_type"`
	Project   string `json:"project"`
}

// ToUrlPathAgg 修改后的转换函数，以TransInfoEntity的trans_type为主
func (t *TransInfoEntity) ToUrlPathAgg() *UrlPathAggEntity {
	urlPathAgg := &UrlPathAggEntity{
		TransType:       t.TransType, // 使用TransInfoEntity的trans_type
		Project:         t.Project,
		ReturnCode:      make(map[string]string),
		ReturnCodeCount: make(map[string]int),
	}

	// 转换ReturnCodes，确保trans_type一致性
	for _, rc := range t.ReturnCodes {
		// 忽略不匹配的trans_type数据
		if rc.TransType != "" && rc.TransType != t.TransType {
			// 可以选择记录日志或跳过
			continue
		}
		urlPathAgg.ReturnCodeCount[rc.ReturnCode] = rc.Count
	}

	return urlPathAgg
}

func (u *UrlPathAggEntity) ToTransInfo() *TransInfoEntity {
	transInfo := &TransInfoEntity{
		Project:   u.Project,
		TransType: u.TransType, // 使用UrlPathAggEntity的url_path作为trans_type
		//TransTypeCn: u.TransTypeCN,
		ReturnCodes: make([]*ReturnCodeEntity, 0, len(u.ReturnCode)),
	}

	// 确保所有ReturnCodeEntity的trans_type与主trans_type一致
	for code, _ := range u.ReturnCode {
		count := 0
		if c, ok := u.ReturnCodeCount[code]; ok {
			count = c
		}

		rcEntity := &ReturnCodeEntity{
			ReturnCode: code,
			Project:    u.Project,
			TransType:  u.TransType, // 强制使用主trans_type
			Status:     "active",
			Count:      count,
		}

		transInfo.ReturnCodes = append(transInfo.ReturnCodes, rcEntity)
	}

	return transInfo
}

func ConvertUrlPathAggListToTransInfoList(urlPathAggList []*UrlPathAggEntity) []*TransInfoEntity {
	// 使用map来去重，以url_path为唯一键
	uniqueMap := make(map[string]*TransInfoEntity)

	for _, upa := range urlPathAggList {
		key := upa.TransType

		if existing, ok := uniqueMap[key]; ok {
			// 合并return codes（通常不应该发生，因为UrlPath应该是唯一的）
			newReturnCodes := upa.ToTransInfo().ReturnCodes
			for _, newRc := range newReturnCodes {
				found := false
				for _, existingRc := range existing.ReturnCodes {
					if existingRc.ReturnCode == newRc.ReturnCode {
						// 更新计数
						existingRc.Count += newRc.Count
						found = true
						break
					}
				}
				if !found {
					existing.ReturnCodes = append(existing.ReturnCodes, newRc)
				}
			}
		} else {
			uniqueMap[key] = upa.ToTransInfo()
		}
	}

	// 转换为数组
	result := make([]*TransInfoEntity, 0, len(uniqueMap))
	for _, ti := range uniqueMap {
		result = append(result, ti)
	}
	return result
}

// 辅助函数：验证TransInfoEntity数据一致性
func (t *TransInfoEntity) ValidateConsistency() []string {
	var errors []string

	for i, rc := range t.ReturnCodes {
		if rc.TransType != "" && rc.TransType != t.TransType {
			errors = append(errors, fmt.Sprintf(
				"ReturnCode[%d]: trans_type不匹配 (期望: %s, 实际: %s)",
				i, t.TransType, rc.TransType,
			))
		}
	}

	return errors
}

// 辅助函数：修复数据一致性
func (t *TransInfoEntity) FixConsistency() {
	for _, rc := range t.ReturnCodes {
		rc.TransType = t.TransType
		if rc.Project == "" {
			rc.Project = t.Project
		}
	}
}

type TransTypeCNEntity struct {
	SysName         string `json:"sys_name"`         // 系统名称代码
	ServiceName     string `json:"service_name"`     // 服务名称代码
	InterfaceEnname string `json:"interface_enname"` // 接口英文名称
	InterfaceName   string `json:"interface_name"`   // 接口中文名称

}
