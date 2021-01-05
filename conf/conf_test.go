package conf_test

import (
	"github.com/spf13/viper"
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
	s.NotNil(viper.Get("run_mode"))

}




// TestViperConfiguration :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(ViperConfigurationSuite))
}
