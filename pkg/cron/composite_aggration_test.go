package cron_test

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// 原始数据结构
type RawData struct {
	RetCode       string `json:"ret_code"`        // 从 biz.expand 中提取
	SelfTest      string `json:"self_test"`       // 从 biz.expand 中提取
	TransChannel  string `json:"trans_channel"`   // 从 biz.expand 中提取
	TransTypeCode string `json:"trans_type_code"` // 从 biz.expand 中提取
	ResultCode    string `json:"result.code"`     // 用于区分成功/失败
}

// 聚合结果结构
type AggregationResult struct {
	RetCode       string `json:"ret_code"`
	SelfTest      string `json:"self_test"`
	TransChannel  string `json:"trans_channel"`
	TransTypeCode string `json:"trans_type_code"`
	Success       int    `json:"success_count"`
	Failure       int    `json:"failure_count"`
}

// 聚合方法
func AggregateData(data []RawData, targetRetCode string) []AggregationResult {
	// 用于存储聚合结果的 map，按维度分组
	resultMap := make(map[string]*AggregationResult)

	// 遍历数据
	for _, d := range data {
		// 创建唯一键值，用于区分不同维度
		key := d.RetCode + "|" + d.TransChannel + "|" + d.TransTypeCode + "|" + string(d.SelfTest)

		// 初始化结果，如果不存在当前维度
		if _, exists := resultMap[key]; !exists {
			resultMap[key] = &AggregationResult{
				RetCode:       d.RetCode,
				SelfTest:      d.SelfTest,
				TransChannel:  d.TransChannel,
				TransTypeCode: d.TransTypeCode,
			}
		}

		// 判断成功/失败
		if d.RetCode == targetRetCode {
			resultMap[key].Success++
		} else {
			resultMap[key].Failure++
		}
	}

	// 将 map 转换为切片返回
	var results []AggregationResult
	for _, r := range resultMap {
		results = append(results, *r)
	}
	return results
}

// 单元测试
func TestAggregateData(t *testing.T) {
	// 准备测试数据
	rawJSON := `
	[
		{"ret_code":"AAAA","self_test":123,"trans_channel":"O999","trans_type_code":"1001","result.code":"201"},
		{"ret_code":"BBBB","self_test":123,"trans_channel":"O999","trans_type_code":"1001","result.code":"500"},
		{"ret_code":"AAAA","self_test":123,"trans_channel":"O999","trans_type_code":"1001","result.code":"201"},
		{"ret_code":"AAAA","self_test":124,"trans_channel":"O999","trans_type_code":"1002","result.code":"200"},
		{"ret_code":"BBBB","self_test":123,"trans_channel":"O999","trans_type_code":"1001","result.code":"404"}
	]`

	var rawData []RawData
	err := json.Unmarshal([]byte(rawJSON), &rawData)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// 定义测试目标值
	targetRetCode := "AAAA"

	// 期望的聚合结果
	expected := []AggregationResult{
		{
			RetCode:       "AAAA",
			SelfTest:      "123",
			TransChannel:  "O999",
			TransTypeCode: "1001",
			Success:       2,
			Failure:       0,
		},
		{
			RetCode:       "BBBB",
			SelfTest:      "123",
			TransChannel:  "O999",
			TransTypeCode: "1001",
			Success:       0,
			Failure:       2,
		},
		{
			RetCode:       "AAAA",
			SelfTest:      "124",
			TransChannel:  "O999",
			TransTypeCode: "1002",
			Success:       1,
			Failure:       0,
		},
	}

	// 调用聚合方法
	results := AggregateData(rawData, targetRetCode)

	// 验证结果是否正确
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Aggregation results do not match expected output.\nGot: %+v\nExpected: %+v", results, expected)
	}
}

// 基准测试
func BenchmarkAggregateData(b *testing.B) {
	// 生成测试数据
	data := generateTestData(10000) // 可以调整数据量，例如 1000, 10000, 100000
	targetRetCode := "AAAA"

	// 重置计时器，避免生成数据的时间影响测试结果
	b.ResetTimer()

	// 多次运行聚合方法
	for i := 0; i < b.N; i++ {
		AggregateData(data, targetRetCode)
	}
}

// 辅助函数：生成测试数据
func generateTestData(size int) []RawData {
	retCodes := []string{"AAAA", "BBBB", "CCCC", "DDDD"}
	transChannels := []string{"O999", "O998", "O997"}
	transTypeCodes := []string{"1001", "1002", "1003"}
	rand.Seed(time.Now().UnixNano())

	data := make([]RawData, size)
	for i := 0; i < size; i++ {
		data[i] = RawData{
			RetCode:       retCodes[rand.Intn(len(retCodes))],
			SelfTest:      "111",
			TransChannel:  transChannels[rand.Intn(len(transChannels))],
			TransTypeCode: transTypeCodes[rand.Intn(len(transTypeCodes))],
			ResultCode:    "200",
		}
	}
	return data
}
