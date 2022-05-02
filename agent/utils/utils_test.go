package utils_test

import (
	"github.com/stretchr/testify/suite"
	utils2 "goOrigin/agent/utils"
	"testing"
)

// HttpClientSuite :
type HttpClientSuite struct {
	suite.Suite
}

func (s *HttpClientSuite) SetupTest() {
}

// TestMarshal :
func (s *HttpClientSuite) TestConfig() {
	utils2.RunShell("")
}

// TestHttpClient :
func TestCmd(t *testing.T) {
	suite.Run(t, new(HttpClientSuite))
}
