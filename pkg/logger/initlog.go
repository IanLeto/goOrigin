package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func NewLogger() *zap.Logger {
	return Logger
}

func InitLogger() {
	// 定义一些需要的变量
	var (
		encoder        zapcore.Encoder
		writer         zapcore.WriteSyncer
		core           zapcore.Core
		consoleEncoder zapcore.Encoder
	)
	// 定义日志配置
	//logConfig := zap.Config{
	//	Level:    zap.AtomicLevel{}, // 日志级别，如 Debug、Info、Warn、Error。AtomicLevel 是一个动态更改日志级别的对象。
	//	Encoding: "json",            // 编码器类型，如 "json"、"console"。
	//	EncoderConfig: zapcore.EncoderConfig{ // 编码器配置
	//		MessageKey:       "MESSAGE", // 日志消息对应的 key，在输出的日志中，消息会以这个 key 输出。
	//		LevelKey:         "LEVEL",   // 日志级别对应的 key，在输出的日志中，级别会以这个 key 输出。
	//		TimeKey:          "TIME",    // 时间戳对应的 key，在输出的日志中，时间戳会以这个 key 输出。
	//		NameKey:          "WAHT?",   // 日志记录器名称对应的 key，在输出的日志中，名称会以这个 key 输出。
	//		CallerKey:        "CALLER",  // 调用者信息对应的 key，在输出的日志中，调用者信息会以这个 key 输出。
	//		FunctionKey:      "FUNC",    // 函数名对应的 key，在输出的日志中，函数名会以这个 key 输出。
	//		StacktraceKey:    "",        // 堆栈信息对应的 key，在输出的日志中，堆栈信息会以这个 key 输出。
	//		SkipLineEnding:   false,     // 是否跳过行结束符。如果为 true，日志的行结束符将不会被写入。
	//		ConsoleSeparator: "",        // 控制台编码器的分隔符。
	//	},
	//	OutputPaths:      nil, // 日志输出路径，可以是任意 io.Writer，如 os.Stdout 或文件路径。
	//	ErrorOutputPaths: nil, // 错误日志输出路径，可以是任意 io.Writer，如 os.Stderr 或文件路径。
	//	InitialFields:    nil, // 初始字段，这些字段将添加到所有日志记录中。
	//}
	loggerConfig := zap.NewProductionEncoderConfig()
	// 设置日志的时间格式为 ISO8601
	loggerConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 使用 JSON 格式的 encoder
	encoder = zapcore.NewJSONEncoder(loggerConfig)
	// 使用控制台的 encoder
	consoleEncoder = zapcore.NewConsoleEncoder(loggerConfig)

	// 打开一个文件用于写日志
	file, _ := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// 将文件作为 zap 的写入器
	writer = zapcore.AddSync(file)

	// 创建一个 zap 的核心
	core = zapcore.NewTee(
		// 第一个核心使用 JSON encoder，并将日志写入文件，日志级别为 Debug
		zapcore.NewCore(encoder, writer, zapcore.DebugLevel),
		// 第二个核心使用控制台 encoder，并将日志写入标准输出，日志级别为 Debug
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))

	// 创建一个 zap logger，并设置 caller（调用者信息）和 stacktrace（堆栈跟踪）的选项
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	// 结束函数
	return
}

func init() {
	InitLogger()
}
