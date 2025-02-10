package processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"goOrigin/internal/model/entity"
	"io"
	"os"
	"sync"
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

	go func() {
		defer close(res) // 确保通道在退出时关闭

		for _, p := range filePath {
			// 打开文件（只读模式）
			file, err := os.Open(p)
			if err != nil {
				fmt.Printf("Failed to open file %s: %v\n", p, err)
				continue
			}
			defer func() { _ = file.Close() }()

			// 创建 bufio.Scanner 按行读取
			scanner := bufio.NewScanner(file)

			// 处理文件内容
			for {
				// 检查是否有 done 信号
				select {
				case <-done:
					return
				default:
					for scanner.Scan() {
						line := scanner.Text() // 读取当前行
						select {
						case res <- line: // 将行内容发送到通道
						case <-done: // 如果收到 done 信号，退出
							return
						}
					}

					// 检查是否遇到错误
					if err := scanner.Err(); err != nil {
						fmt.Printf("Error reading file %s: %v\n", p, err)
						break
					}

					// 如果到达文件末尾，暂停一段时间等待文件写入新内容
					time.Sleep(100 * time.Millisecond)
				}

				// 如果有新行，发送到通道

			}
		}
	}()

	return res
}

var AggData = func(done <-chan interface{}, data <-chan []byte, condition func(a any) any) <-chan []byte {
	res := make(chan []byte)
	var (
		err error
		wg  sync.WaitGroup
	)

	return res
}
