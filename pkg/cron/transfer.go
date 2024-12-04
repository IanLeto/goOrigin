package cron

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goOrigin/config"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
	"goOrigin/pkg/processor"
	"io"
	"os"
	"time"
)

// 初始化 zap.Logger
var logger = func() *zap.Logger {
	// 配置 zapcore.Encoder (日志格式配置)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                        // 时间字段的键名
		LevelKey:       "level",                       // 日志级别字段的键名
		NameKey:        "logger",                      // 日志器名称字段的键名
		CallerKey:      "caller",                      // 调用者信息字段的键名
		MessageKey:     "msg",                         // 日志消息字段的键名
		StacktraceKey:  "stacktrace",                  // 堆栈跟踪字段的键名
		LineEnding:     zapcore.DefaultLineEnding,     // 每行日志的结束符
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 日志级别大写编码 (INFO, WARN, ERROR)
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // 时间格式 (ISO8601 格式)
		EncodeDuration: zapcore.StringDurationEncoder, // 时间间隔格式
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 调用者信息格式 (简短路径)
	}

	// 设置日志输出级别 (DEBUG、INFO、WARN、ERROR)
	level := zapcore.DebugLevel // 写死为 DEBUG 级别，可以用配置加载

	// 日志写入目标
	// 1. 输出到标准输出 (控制台)
	consoleWriter := zapcore.Lock(os.Stdout)

	// 2. 输出到文件 (额外配置)
	fileWriter, err := os.Create("app.log") // 写死日志文件路径为 "app.log"
	if err != nil {
		panic("无法创建日志文件: " + err.Error())
	}

	// 创建 zapcore.Core
	core := zapcore.NewTee(
		// 使用 Tee 将日志同时写入多个输出
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriter, level),            // 控制台输出
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(fileWriter), level), // 文件输出 (JSON 格式)
	)

	// 构建 zap.Logger 实例
	logger := zap.New(core,
		zap.AddCaller(),                       // 显示调用者信息 (文件名和行号)
		zap.AddCallerSkip(1),                  // 调用栈跳过级别 (适配封装的 logger)
		zap.AddStacktrace(zapcore.ErrorLevel), // 仅在 ERROR 级别记录堆栈
	)

	return logger
}()

type Transfer struct {
	*entity.Record
	Alias string
}

func (t *Transfer) Exec(ctx context.Context) error {
	var (
		//esClient = elastic.GlobalEsConns[config.ConfV2.Base.Region]
		//index    = t.Alias
		err error
	)
	file, err := os.Open("./test.json")
	defer func() { _ = file.Close() }()

	// 创建一个缓冲区读取器,用于读取 JSON 数据
	reader := bufio.NewReader(file)

	// 创建一个字节缓冲区,用于存储读取的数据
	var buffer bytes.Buffer

	// 循环读取 JSON 数据,直到读取完毕或出错
	for {
		// 读取一行数据
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				// 读取完毕,跳出循环
				break
			}
			logrus.Errorf("Failed to read record data: %v", err)
			return err
		}

		// 将读取的数据写入字节缓冲区
		buffer.Write(line)
	}

	// 将字节缓冲区中的数据转换为字节数组
	body := buffer.Bytes()
	logrus.Infof("Read record data: %s", body)
	// 调用 EsV2Conn 的 Create 方法写入数据到 Elasticsearch
	//resp, err := esClient.Create(index, body)

	if err != nil {
		logrus.Errorf("Failed to create document in Elasticsearch: %v", err)
		return err
	}

	// 处理 Elasticsearch 的响应
	//logrus.Infof("Document created in Elasticsearch. Response: %s", resp)

	return nil
}

func (t *Transfer) Exec2(ctx context.Context) error {
	var (
		esClient = elastic.GlobalEsConns[config.ConfV2.Base.Region]
		index    = t.Alias
		err      error
	)

	// 将Record结构体转换为JSON格式的字节数组
	body, err := json.Marshal(t.Record)
	if err != nil {
		logrus.Errorf("Failed to marshal record: %v", err)
		return err
	}

	// 调用EsV2Conn的Create方法写入数据到Elasticsearch
	resp, err := esClient.Create(index, body)
	if err != nil {
		logrus.Errorf("Failed to create document in Elasticsearch: %v", err)
		return err
	}

	// 处理Elasticsearch的响应
	logrus.Infof("Document created in Elasticsearch. Response: %s", resp)

	return nil
}

func (t *Transfer) GetName() string {
	//TODO implement me
	return ""
}

// TransferCornFactory : 最重要的方法，如何注册并启动任务
func TransferCornFactory() error {
	for {
		select {
		case <-time.NewTicker(time.Duration(10) * time.Second).C:

			GTM.AddJob(&Transfer{})

		}
	}
	return nil
}
func TransferCornFactory2() error {
	var (
		interval int
		//err      error
	)
	transferConf := config.ConfV2.Env[config.ConfV2.Base.Region].CronJobConfig.TransferConfig
	interval = transferConf.Interval

	for {
		select {
		case <-time.NewTicker(time.Duration(interval) * time.Second).C:
			var (
				recordTables = make([]*dao.TRecord, 0)
				err          error
			)
			db := mysql.GlobalMySQLConns[config.ConfV2.Base.Region]
			sql := db.Client.Table("t_records")
			sql.Where("created_at > ?", time.Now().Add(-time.Minute))
			tRecords := sql.Find(&recordTables)
			if tRecords.Error != nil {
				logrus.Errorf("create record failed %s: %s", err, tRecords.Error)
				continue
			}
			logger.Sugar().Infoln("根据每条数据的特性，开始封装数据流水线")
			func() {
				// 读取配置文件，封装数据流水线
				var pipeline = processor.Pipeline{}
				var node1 = &processor.FilterProcessor{}
				pipeline.Add([]processor.Processor{node1})

			}()

			for _, recordTable := range recordTables {
				GTM.AddJob(&Transfer{
					Record: repository.ToRecordEntity(recordTable),
					Alias:  transferConf.ES.Alias,
				})
			}

		}
	}

	return nil
}
