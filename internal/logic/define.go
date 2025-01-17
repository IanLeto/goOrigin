package logic

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

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
