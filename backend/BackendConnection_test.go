package backend_test

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/suite"
	"goOrigin/backend"
	"goOrigin/config"
	"testing"
)

// ConnectionConfigurationSuite :
type ConnectionConfigurationSuite struct {
	suite.Suite
	client *backend.MySQLBackend
}

func (s *ConnectionConfigurationSuite) SetupTest() {
	var err error
	s.NoError(config.InitConfig())
	s.client, err = backend.NewMySQLBackend("")
	s.NoError(err)
}

// TestMarshal :
func (s *ConnectionConfigurationSuite) TestConfig() {
	s.NoError(s.client.Client.DB().Ping())
}

// TestViperConfiguration :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(ConnectionConfigurationSuite))
}
