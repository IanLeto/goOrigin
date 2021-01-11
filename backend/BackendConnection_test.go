package backend_test

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/suite"
	"goOrigin/backend"
	"goOrigin/conf"
	"testing"
)

// ConnectionConfigurationSuite :
type ConnectionConfigurationSuite struct {
	suite.Suite
	db *gorm.DB
}

func (s *ConnectionConfigurationSuite) SetupTest() {
	var err error
	s.NoError(conf.InitConfig())
	s.db, err = backend.NewMySQLBackend("")
	s.NoError(err)
}

// TestMarshal :
func (s *ConnectionConfigurationSuite) TestConfig() {
	s.NoError(s.db.DB().Ping())
}

// TestViperConfiguration :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(ConnectionConfigurationSuite))
}
