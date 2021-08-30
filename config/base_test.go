package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// ViperConfigurationSuite :
type ViperConfigurationSuite struct {
	suite.Suite
}

func (s *ViperConfigurationSuite) SetupTest() {
}

// TestMarshal :
func (s *ViperConfigurationSuite) TestConfig() {

}

func (s ViperConfigurationSuite) TestMySqlBackendConfig() {

	//s.Equal("localhost:3306", config..Backend.MySqlBackendConfig.Address)

}

// TestViperConfiguration :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(ViperConfigurationSuite))
}
