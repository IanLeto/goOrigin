package processor_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/processor"
	"math/rand"
	"sync"
	"testing"
)

// 定义一个对象池，用于复用 KafkaLogEntity 实例
var kafkaLogEntityPool = sync.Pool{
	New: func() interface{} {
		return &entity.KafkaLogEntity{} // 创建新的 KafkaLogEntity 实例
	},
}

// 单条数据处理函数（使用对象池）
var DataConvWithPool = func(value []byte) ([]byte, error) {
	// 从对象池获取 KafkaLogEntity 实例
	logEntity := kafkaLogEntityPool.Get().(*entity.KafkaLogEntity)

	// 确保用完后将对象放回池中
	defer func() {
		// 清空结构体内容，避免残留数据
		*logEntity = entity.KafkaLogEntity{}
		kafkaLogEntityPool.Put(logEntity)
	}()

	err := json.Unmarshal(value, logEntity)
	if err != nil {
		return nil, err
	}
	if logEntity.Trans.RetCode != "0000" {
		logEntity.TraceId = "11"
	}
	return json.Marshal(logEntity)
}

// 基于 channel 的批量数据处理函数（使用对象池）
var DataConvStageWithPool = func(value chan []byte, workers int) chan []byte {
	res := make(chan []byte)

	// 启动指定数量的 worker goroutines
	go func() {
		defer close(res)
		sem := make(chan struct{}, workers) // 控制并发度
		var wg sync.WaitGroup               // 确保所有 goroutines 完成
		for i := range value {
			wg.Add(1)
			sem <- struct{}{} // 占用一个 worker
			go func(data []byte) {
				defer func() {
					<-sem // 释放 worker
					wg.Done()
				}()

				// 从对象池获取 KafkaLogEntity 实例
				logEntity := kafkaLogEntityPool.Get().(*entity.KafkaLogEntity)

				// 确保用完后将对象放回池中
				defer func() {
					*logEntity = entity.KafkaLogEntity{}
					kafkaLogEntityPool.Put(logEntity)
				}()

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
		wg.Wait() // 等待所有 goroutines 完成
	}()

	return res
}

// 辅助函数：生成随机 Kafka 日志数据
func generateTestData(size int) [][]byte {
	data := make([][]byte, size)
	for i := 0; i < size; i++ {
		entity := &entity.KafkaLogEntity{}
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

// Benchmark for DataConvWithPool
func BenchmarkDataConvWithPool(b *testing.B) {
	data := generateTestData(100000) // 生成 100,000 条测试数据

	b.ResetTimer() // 重置计时器
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			_, _ = DataConvWithPool(d)
		}
	}
}

// Benchmark for DataConvStageWithPool
func BenchmarkDataConvStageWithPool(b *testing.B) {
	data := generateTestData(100000) // 生成 100,000 条测试数据

	b.ResetTimer() // 重置计时器
	for i := 0; i < b.N; i++ {
		input := make(chan []byte, len(data))
		for _, d := range data {
			input <- d
		}
		close(input)

		output := DataConvStageWithPool(input, 30) // 使用 10 个 worker
		for range output {
			// 消费结果
		}
	}
}

type filePathPipe struct {
	filePath []string
}

// StageSuite :
type StageSuite struct {
	suite.Suite
	filePathPipe *filePathPipe
}

func (s *StageSuite) SetupTest() {
	s.filePathPipe = &filePathPipe{}
	s.filePathPipe.filePath = []string{"/home/ian/workdir/goOrigin/pkg/processor/span.log"}
}

// TestMarshal :
func (s *StageSuite) TestConfig() {
	var ch = make(chan struct{})
	for v := range processor.FileReadHead(ch, s.filePathPipe.filePath[0]) {
		fmt.Println(v)
	}

}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(StageSuite))
}
