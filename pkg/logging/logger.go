package logging

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"goOrigin/config"
	"log"
	"sync"
	"time"
)

var StdLogger *Logger
var NilLogger *Logger

func GetStdLogger() *Logger {
	return StdLogger
}

// Logger 我们自定义的log 对象
type Logger struct {
	configOnce sync.Once
	*logrus.Logger
}

func newHook() logrus.Hook {
	writer, err := rotatelogs.New("config"+config.Conf.Logging.FileName,
		// 为日志建立连接
		rotatelogs.WithLinkName("common"),
		// 每24小时分割一次
		rotatelogs.WithRotationTime(time.Duration(config.Conf.Logging.Rotation.Count)*time.Hour),
		// 文件最多保留个数
		rotatelogs.WithRotationCount(cast.ToUint(config.Conf.Logging.Rotation.Count)),
	)
	if err != nil {
		panic(err)
	}

	errWriter, err := rotatelogs.New("error"+config.Conf.Logging.FileName,
		// 为日志建立连接
		rotatelogs.WithLinkName("error"),
		// 每24小时分割一次
		rotatelogs.WithRotationTime(time.Duration(config.Conf.Logging.Rotation.Count)*time.Hour),
		// 文件最多保留个数
		rotatelogs.WithRotationCount(cast.ToUint(config.Conf.Logging.Rotation.Count)),
	)
	if err != nil {
		panic(err)
	}
	divHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: errWriter,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})
	return divHook
}

func InitLogger() error {
	logger := logrus.StandardLogger()
	StdLogger = &Logger{
		Logger: logger,
	}
	log.SetOutput(StdLogger.Writer()) // 将基础log 中的文件输出，定向到logrus
	logger = logrus.New()
	logrus.AddHook(newHook())
	logger.Level = logrus.PanicLevel
	NilLogger = &Logger{
		Logger: logger,
	}
	return nil
}
