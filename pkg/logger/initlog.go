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
	var (
		encoder        zapcore.Encoder
		writer         zapcore.WriteSyncer
		core           zapcore.Core
		consoleEncoder zapcore.Encoder
	)
	loggerConfig := zap.NewProductionEncoderConfig()
	loggerConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder = zapcore.NewJSONEncoder(loggerConfig)
	consoleEncoder = zapcore.NewConsoleEncoder(loggerConfig)

	file, _ := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer = zapcore.AddSync(file)

	core = zapcore.NewTee(zapcore.NewCore(encoder, writer, zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return
}

func init() {
	InitLogger()
}
