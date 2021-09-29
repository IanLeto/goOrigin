package shell_test

import (
	"fmt"
	"goOrigin/config"
	"goOrigin/pkg/shell"
	"goOrigin/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

// SSHClientSuite :
type SSHClientSuite struct {
	suite.Suite
	conf *config.Config
	conn *shell.SSHSession
}

func (s *SSHClientSuite) SetupTest() {
	config.InitConf(utils.GetFilePath("config.yaml"))
	shell.InitSSH()
	s.conn = shell.SSHConn
}

// TestMarshal :
func (s *SSHClientSuite) TestConfig() {
	res, err := s.conn.Exec("pwd")
	s.NoError(err)
	fmt.Println(string(res))
}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(SSHClientSuite))
}
