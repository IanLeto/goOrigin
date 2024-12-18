package cron

import (
	"bufio"
	"context"
	"fmt"
	"goOrigin/config"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/processor"
	"os"
	"os/signal"
	"syscall"
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
		file    *os.File
		err     error
		signals = make(chan os.Signal, 1)
	)
	var (
		nodes = make([]processor.Node, 0)
		pipe  = processor.Pipeline{}
	)
	//nodes = append(nodes, &processor.MetricProcessor{RetCodes: []string{"AAAAA"}})
	pipe.Nodes = nodes
	// 打开日志文件
	file, err = os.Open(c.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer func() { _ = file.Close() }()

	// 捕获终止信号
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// 创建一个带偏移量的缓冲读取器
	reader := bufio.NewScanner(file)

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
					ch <- reader.Bytes()
					// 处理日志行
					pipe.Start(ctx, ch)
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
