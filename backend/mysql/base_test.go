package mysql_test

import (
	"github.com/stretchr/testify/suite"
	"goOrigin/backend/mysql"
	"goOrigin/config"
	"testing"
)

// HttpClientSuite :
type MysqlBackendSuite struct {
	suite.Suite
	conf    *config.Config
	backend *mysql.MySQLBackend
}

func (s *MysqlBackendSuite) SetupTest() {
	s.conf = config.GlobalConfig

}

// TestMarshal :
func (s *MysqlBackendSuite) TestConfig() {
	if s.conf.RunMode != "debug" {
		return
	}
	s.backend = mysql.MySqlBackend
	s.NoError(s.backend.Client.DB().Ping())

}

// TestMysqlBackend :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(MysqlBackendSuite))
}
