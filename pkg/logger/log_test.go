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

// TestHttpClient :
func TestLogConfiguration(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}
