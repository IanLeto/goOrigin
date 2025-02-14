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

// FileReadHead 监听文件新增内容，并通过通道返回新行
var FileReadHead = func(done <-chan interface{}, filePath ...string) <-chan string {
	res := make(chan string) // 输出通道，用于发送读取到的行内容

	for _, p := range filePath {
		// 打开文件（只读模式）
		go func(path string) {
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Failed to open file %s: %v\n", path, err)
				return
			}
			defer file.Close()

			// 获取文件当前大小，跳转到末尾
			stat, err := file.Stat()
			if err != nil {
				fmt.Printf("Failed to get file info %s: %v\n", path, err)
				return
			}
			offset := stat.Size()
			file.Seek(offset, 0) // 从文件末尾开始监听

			reader := bufio.NewReader(file)

			for {
				select {
				case <-done:
					return // 如果收到 `done` 信号，退出
				default:
					// 读取新内容
					line, err := reader.ReadString('\n')
					if err != nil { // 说明暂时没有新内容
						// **检测文件是否增长**
						newStat, err := file.Stat()
						if err == nil && newStat.Size() > offset {
							// **文件变大，调整偏移量**
							offset, _ = file.Seek(0, os.SEEK_CUR)
						}
						time.Sleep(100 * time.Millisecond) // 没有新内容，等待一会儿
						continue
					}

					// 发送新内容到通道
					select {
					case res <- line:
					case <-done:
						return
					}

					// 更新偏移量
					offset, _ = file.Seek(0, os.SEEK_CUR)
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
