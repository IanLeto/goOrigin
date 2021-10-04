package logging

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"goOrigin/config"
	"goOrigin/pkg/utils"
	"sync"
	"time"
)

var StdLogger *Logger

func GetStdLogger() *Logger {
	return StdLogger
}

// Logger 我们自定义的log 对象
type Logger struct {
	configOnce sync.Once
	*logrus.Logger
}

func InitLogging() error {
	logrus.AddHook(newHook())
	switch config.Conf.RunMode {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "init":
		logrus.SetLevel(logrus.DebugLevel)
	case "":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
	return nil
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
	utils.NoError(err)

	errWriter, err := rotatelogs.New("error"+config.Conf.Logging.FileName,
		// 为日志建立连接
		rotatelogs.WithLinkName("error"),
		// 每24小时分割一次
		rotatelogs.WithRotationTime(time.Duration(config.Conf.Logging.Rotation.Count)*time.Hour),
		// 文件最多保留个数
		rotatelogs.WithRotationCount(cast.ToUint(config.Conf.Logging.Rotation.Count)),
	)
	utils.NoError(err)
	divHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: errWriter,
		logrus.FatalLevel: writer, // todo
		logrus.PanicLevel: writer, // todo
	}, &logrus.TextFormatter{DisableColors: true})
	return divHook
}

func init() {

}
