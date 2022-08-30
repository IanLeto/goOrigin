package DB_test

import (
	"github.com/stretchr/testify/suite"
	"goOrigin/backend/DB"
	"goOrigin/config"
	"testing"
)

// HttpClientSuite :
type MysqlBackendSuite struct {
	suite.Suite
	conf    *config.Config
	backend *DB.MySQLBackend
}

func (s *MysqlBackendSuite) SetupTest() {
	s.conf = nil

}

// TestMarshal :
func (s *MysqlBackendSuite) TestConfig() {
	if s.conf.RunMode != "debug" {
		return
	}
	s.backend = DB.MySqlBackend
	s.NoError(s.backend.Client.DB().Ping())

}

// TestMysqlBackend :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(MysqlBackendSuite))
}
