package processor

import (
	"encoding/json"
	"goOrigin/internal/model/entity"
	"io"
	"os"
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

var FileReadHead = func(done <-chan interface{}, filePath ...string) <-chan []byte {
	var (
		file *os.File
		res  = make(chan []byte)
	)
	go func() {
		defer close(res)
		for _, p := range filePath {
			select {
			case <-done:
				return
			default:
				file, err = os.OpenFile(p, os.O_CREATE, 0644)
				value, err := io.ReadAll(file)
				logger.Sugar().Infof("file value: %s", string(value))
				if err != nil {
					logger.Sugar().Errorf("file value: %s", err)
				}
				res <- value
			}
		}
	}()
	return res
}
