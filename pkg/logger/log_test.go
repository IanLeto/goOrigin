package logger_test

import (
	"github.com/stretchr/testify/suite"
	"goOrigin/config"
	"goOrigin/pkg/logger"
	"testing"
)

type LoggerSuite struct {
	suite.Suite
	conf *config.Config
}

func (s *LoggerSuite) SetupTest() {
	logger.InitLogger()

}

// TestMarshal :
func (s *LoggerSuite) TestConfig() {
	logger.Logger.Info("test")
	// https://darjun.github.io/2020/04/23/godailylib/zap/
}

// TestHttpClient :
func TestLogConfiguration(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}
