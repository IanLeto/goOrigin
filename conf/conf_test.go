package conf_test

import (
	"goOrigin/conf"
	"testing"

	"github.com/stretchr/testify/suite"
)

// ViperConfigurationSuite :
type ViperConfigurationSuite struct {
	suite.Suite
}

func (v *ViperConfigurationSuite) SetupTest() {
	conf.InitConfig()
}

// TestMarshal :
func (s *ViperConfigurationSuite) TestConfig() {


}

func (s ViperConfigurationSuite) TestMySqlBackendConfig() {

	s.Equal("localhost", conf.Conf.Backend.MySqlBackendConfig.Address)

}

// TestViperConfiguration :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(ViperConfigurationSuite))
}
