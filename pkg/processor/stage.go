package processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"goOrigin/internal/model/entity"
	"io"
	"os"
	"time"
)

// FileRead Stage is a stage in the pipeline
var FileRead = func(filePath string) ([]byte, error) {
	var (
		file *os.File
		err  error
	)

	file, err = os.OpenFile(filePath, os.O_CREATE, 0644)
	defer func() { _ = file.Close() }()
	value, err := io.ReadAll(file)
	logger.Sugar().Infof("file value: %s", string(value))
	return value, err
}

// FileRead Stage is a stage in the pipeline
var FileReadStage = func(filePath <-chan string) (chan string, error) {
	var (
		file *os.File
		err  error
		res  chan string
	)

	go func() {
		defer close(res)
		defer func() { _ = file.Close() }()
		for p := range filePath {
			file, err = os.OpenFile(p, os.O_CREATE, 0644)
			value, err := io.ReadAll(file)
			logger.Sugar().Infof("file value: %s", string(value))
			if err != nil {
				logger.Sugar().Errorf("file value: %s", err)
			}
			res <- string(value)
		}
	}()

	return res, err
}

var DataConv = func(value []byte) ([]byte, error) {
	var (
		logEntity *entity.KafkaLogEntity
		err       error
	)
	err = json.Unmarshal(value, logEntity)
	if err != nil {
		logger.Sugar().Errorf("file value: %s", err)
		return nil, err
	}
	if logEntity.Trans.RetCode != "0000" {
		logEntity.TraceId = "11"
	}
	value, err = json.Marshal(logEntity)
	return value, err

}

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
				logEntity := &entity.KafkaLogEntity{}
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

func FileWrite(filePath string, value []byte) ([]byte, error) {
	var (
		file *os.File
		err  error
	)

	file, err = os.OpenFile(filePath, os.O_CREATE, 0644)
	defer func() { _ = file.Close() }()
	_, err = file.Write([]byte(value))
	return value, err
}

var FileReadHead = func(done <-chan interface{}, filePath ...string) <-chan string {
	res := make(chan string) // 输出通道，用于发送读取到的行内容

	for _, p := range filePath {
		go func(path string) {
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Failed to open file %s: %v\n", path, err)
				return
			}
			defer file.Close()

			// 创建一个 Scanner 逐行读取文件
			reader := bufio.NewScanner(file)

			// 读取文件已有的内容（从头开始）
			for reader.Scan() {
				select {
				case res <- reader.Text(): // 发送读取的行内容
				case <-done:
					close(res)
					return
				}
			}

			// 检查是否有错误
			if err := reader.Err(); err != nil {
				fmt.Printf("Error reading file %s: %v\n", path, err)
				return
			}

			// 监听文件新增内容
			for {
				select {
				case <-done:
					close(res)
					return
				default:
					// 读取新内容
					if reader.Scan() {
						res <- reader.Text()
					} else {
						time.Sleep(500 * time.Millisecond) // 休眠一段时间，避免 CPU 过载
					}
				}
			}
		}(p)
	}

	return res
}

var AggData = func(done <-chan interface{}, data <-chan []byte, condition func(a any) any) <-chan []byte {
	res := make(chan []byte)
	var (
	//err     error
	//wg      sync.WaitGroup
	//buckets = make([]struct{}, 50)
	)

	return res
}

var AggDataStage = func(done <-chan interface{}, data <-chan []byte, condition func(a any) any, workers int) <-chan []byte {
	panic(1)
	//var (
	//	out = make(chan []byte)
	//	wg  sync.WaitGroup
	//)
	//for i := 0; i < workers; i++ {
	//	wg.add(1)
	//}
}
