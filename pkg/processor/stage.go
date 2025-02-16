package processor

import (
	"encoding/json"
	"fmt"
	"github.com/hpcloud/tail"
	_ "github.com/hpcloud/tail"
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

// FileReadHead 监听多个文件的新增内容，并通过通道返回新行
func FileReadHead(done <-chan struct{}, filePaths ...string) <-chan string {
	res := make(chan string) // 输出通道

	for _, filePath := range filePaths {
		go func(path string) {
			config := tail.Config{
				ReOpen:    true,                                 // 文件被移动或删除后重新打开
				Follow:    true,                                 // 持续监听新内容
				Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件末尾开始读取
				MustExist: false,                                // 文件不存在时不报错，等待创建
				Poll:      true,                                 // 轮询模式，适用于 inotify 不支持的平台
			}

			tails, err := tail.TailFile(path, config)
			if err != nil {
				fmt.Printf("Failed to tail file %s: %v\n", path, err)
				return
			}
			defer tails.Cleanup() // 退出时释放资源

			// 监听文件内容
			for {
				select {
				case <-done: // 监听到退出信号，退出 goroutine
					fmt.Printf("Stopping file tailing: %s\n", path)
					return
				case line, ok := <-tails.Lines:
					if !ok {
						fmt.Printf("Tail file closed, reopening: %s\n", path)
						time.Sleep(time.Second) // 等待 1 秒后重新打开
						continue
					}
					res <- line.Text
				}
			}
		}(filePath)
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
