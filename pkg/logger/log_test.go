package logger_test

import (
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"goOrigin/config"
	"goOrigin/pkg/logger"
	"testing"
)

type LoggerSuite struct {
	suite.Suite
	conf *config.Config
	log  *zap.Logger
}

func (s *LoggerSuite) SetupTest() {
	s.log, _ = logger.InitZap()
}

// TestMarshal : 基础输出
func (s *LoggerSuite) TestConfig() {
	suger := s.log.Sugar()
	suger.Info("基础用法")
}

// TestMarshal : 日志等级以及自定义config
func (s *LoggerSuite) TestConfig2() {
	var (
		err error
	)
	conf := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.ErrorLevel), // Set to ErrorLevel
		Encoding:         "json",                               // 必填项
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),    // 必填项
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	s.log, err = conf.Build()
	s.NoError(err)
	s.log.Sugar().Info("不会打印")
	s.log.Sugar().Error("会打印")
}

// TestMarshal :
func (s *LoggerSuite) TestStack() {
	var (
		err error
	)
	conf := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.ErrorLevel), // Set to ErrorLevel
		Encoding:         "json",                               // 必填项
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),    // 必填项
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	s.log, err = conf.Build()
	s.NoError(err)
	s.log.Sugar().Info("不会打印")
	s.log.Sugar().Error("会打印")
}

// TestHttpClient :
func TestLogConfiguration(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}
