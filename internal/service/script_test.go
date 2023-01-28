package service

import (
	"github.com/stretchr/testify/suite"
	"goOrigin/config"
	"testing"
)

type ScriptAPISuite struct {
	suite.Suite
	conf *config.Config
}

func (s *ScriptAPISuite) SetupTest() {

}

// TestMarshal :
func (s *ScriptAPISuite) TestConfig() {
}

// TestHttpClient :
func TestScriptConfiguration(t *testing.T) {
	suite.Run(t, new(ScriptAPISuite))
}
