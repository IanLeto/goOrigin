package processor_test

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"goOrigin/pkg/processor"
	"math/rand"
	"testing"
)

// 定义 Kafka 日志实体
type KafkaLogEntity struct {
	Trans struct {
		RetCode string `json:"ret_code"`
	} `json:"trans"`
	TraceId string `json:"trace_id"`
}

// 单条数据处理函数
var DataConv = func(value []byte) ([]byte, error) {
	// 初始化 logEntity
	logEntity := &KafkaLogEntity{}
	err := json.Unmarshal(value, logEntity)
	if err != nil {
		return nil, err
	}
	if logEntity.Trans.RetCode != "0000" {
		logEntity.TraceId = "11"
	}
	return json.Marshal(logEntity)
}

// 基于 channel 的批量数据处理函数
var DataConvStage = func(value chan []byte, workers int) chan []byte {
	res := make(chan []byte)

	// 启动指定数量的 worker goroutines
	go func() {
		defer close(res)
		sem := make(chan struct{}, workers) // 控制并发度
		for i := range value {
			sem <- struct{}{} // 占用一个 worker
			go func(data []byte) {
				defer func() { <-sem }() // 释放 worker
				logEntity := &KafkaLogEntity{}
				err := json.Unmarshal(data, logEntity)
				if err != nil {
					return
				}
				if logEntity.Trans.RetCode != "0000" {
					logEntity.TraceId = "11"
				}
				ephValue, err := json.Marshal(logEntity)
				if err != nil {
					return
				}
				res <- ephValue
			}(i)
		}
		// 等待所有 goroutines 结束
		for i := 0; i < cap(sem); i++ {
			sem <- struct{}{}
		}
	}()

	return res
}

// 辅助函数：生成随机 Kafka 日志数据
func generateTestData(size int) [][]byte {
	data := make([][]byte, size)
	for i := 0; i < size; i++ {
		entity := KafkaLogEntity{}
		if rand.Intn(2) == 0 {
			entity.Trans.RetCode = "0000"
		} else {
			entity.Trans.RetCode = "1234"
		}
		value, _ := json.Marshal(entity)
		data[i] = value
	}
	return data
}

// 基准测试：测试 DataConv
func BenchmarkDataConv(b *testing.B) {
	data := generateTestData(10000) // 生成 10,000 条测试数据
	b.ResetTimer()                  // 重置计时器，忽略数据生成时间
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			_, _ = DataConv(d)
		}
	}
}

// 基准测试：测试 DataConvStage
func BenchmarkDataConvStage(b *testing.B) {
	data := generateTestData(10000) // 生成 10,000 条测试数据
	input := make(chan []byte, len(data))
	for _, d := range data {
		input <- d
	}
	close(input)

	b.ResetTimer() // 重置计时器，忽略数据生成时间
	for i := 0; i < b.N; i++ {
		output := DataConvStage(input, 1) // 使用 10 个 worker
		for range output {
			// 消费结果
		}
	}
}

type StageSuite struct {
	suite.Suite
}

func (s *StageSuite) SetupTest() {

}

// TestMarshal :
func (s *StageSuite) TestConfig() {
	v, err := processor.FileRead("/home/ian/workdir/goOrigin/pkg/processor/test.txt")
	s.NoError(err)
	s.Equal("1", v)

}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(StageSuite))
}
