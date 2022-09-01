package logger

import (
	"go.uber.org/zap"
)

//var Logger *zap.SugaredLogger
var Logger *zap.Logger

func InitLogger() {
	Logger, _ = zap.NewProduction()
	//Logger = logger.Sugar()
	//var (
	//	encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//	writeSyncer =
	//)
	//
	//logger, _ = zap.New()
}

func init() {
	InitLogger()

}
