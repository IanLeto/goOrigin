package httpclient_test

import (
	"goOrigin/config"
	"goOrigin/httpclient"
	"testing"

	"github.com/stretchr/testify/suite"
)

// HttpClientSuite :
type HttpClientSuite struct {
	suite.Suite
	conf *config.Config
}

func (s *HttpClientSuite) SetupTest() {
	s.conf = config.GlobalConfig

}

// TestMarshal :
func (s *HttpClientSuite) TestConfig() {
	if s.conf.RunMode != "debug" {
		return
	}
	var ch = make(chan struct{})
	client := httpclient.NewCCClient(nil)
	go func() {
		s.NoError(httpclient.PingCCClient(client, ch)())
	}()
	<-ch

}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(HttpClientSuite))
}
