package cron

import (
	"bufio"
	"context"
	"fmt"
	"goOrigin/config"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/processor"
	"os"
	"path/filepath"
	"time"
)

type Consumer struct {
	*entity.ODAMetricEntity
	FilePath string
}

// processLine 处理每一行日志数据
func (c *Consumer) processLine(line string) {
	// 在这里实现你的日志行处理逻辑
	fmt.Printf("Processing line: %s\n", line)
}

// Exec 按行读取日志文件，并实时处理新增的日志内容
func (c *Consumer) Exec(ctx context.Context) error {
	var (
		file *os.File
		err  error
		//signals = make(chan os.Signal, 1)
	)
	var (
		nodes = make([]processor.Node, 0)
		pipe  = processor.Pipeline{}
	)
	pipe.Nodes = nodes
	// 打开日志文件
	file, err = os.Open(c.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer func() { _ = file.Close() }()
	actualPath, err := filepath.Abs(c.FilePath)
	if err != nil {
		return fmt.Errorf("获取实际文件路径时出错: %v", err)
	}

	// 打印实际读取的文件路径
	logger.Sugar().Infof("实际读取的文件路径: %s", actualPath)
	// 捕获终止信号
	file, err = os.Open(actualPath)
	// 创建一个带偏移量的缓冲读取器
	reader := bufio.NewScanner(file)
	//for reader.Scan() {
	//	logger.Sugar().Info("实际内容", reader.Text())
	//	//ch <- reader.Bytes()
	//	// 处理日志行
	//	//pipe.Start(ctx, ch)
	//}
	if err := reader.Err(); err != nil {
		fmt.Println("读取文件时发生错误:", err)
	}
	// 启动一个 Goroutine 来处理新增的日志内容
	go func() {
		var ch = make(chan []byte)
		defer func() { close(ch) }()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context canceled, stopping log reading...")
				return
			default:
				// 检查是否有新行可读
				for reader.Scan() {
					logger.Sugar().Info("实际内容", reader.Text())
					//ch <- reader.Bytes()
					// 处理日志行
					//pipe.Start(ctx, ch)
				}

				// 如果读取完成但没有新内容，等待文件更新
				time.Sleep(500 * time.Millisecond) // 可根据需要调整刷新频率
			}
		}
	}()

	// 等待退出信号
	//<-signals
	//fmt.Println("Received termination signal, shutting down...")
	return nil
}

func (c Consumer) GetName() string {
	return ""
}
func ConsumerFactory() error {
	var ()
	var (
		consumer = config.ConfV2.Env[config.ConfV2.Base.Region].CronJobConfig.TransferConfig
	)
	for {
		select {
		case <-time.NewTicker(time.Duration(consumer.Interval) * time.Second).C:

			GTM.AddJob(&Consumer{
				FilePath: "./test.json",
			})

		}
	}
	return nil
}
