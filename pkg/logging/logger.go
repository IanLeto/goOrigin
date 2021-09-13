package logging

import (
	"github.com/heirko/go-contrib/logrusHelper"
	"github.com/heralight/logrus_mate"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
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

func (l *Logger) InitConfig() error {
	l.configOnce.Do(func() {
		logrusHelper.SetConfig(l.Logger, logrus_mate.LoggerConfig{

		})
	})
	return nil
}

func init() {
	// logrus 实例
	logger := logrus.StandardLogger()

	StdLogger = &Logger{
		Logger: logger,
	}
	log.SetOutput(StdLogger.Writer()) // 将基础log 中的文件输出，定向到logrus

	logger = logrus.New()

	logger.Level = logrus.PanicLevel
	NilLogger = &Logger{
		Logger: logger,
	}
}
