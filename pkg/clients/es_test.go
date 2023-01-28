package clients_test

import (
	"github.com/stretchr/testify/suite"
	"goOrigin/pkg/clients"
	"testing"
)

type EsAPISuite struct {
	suite.Suite
	conn *clients.EsConn
}

func (s *EsAPISuite) SetupTest() {

}

// TestMarshal :
func (s *EsAPISuite) TestConfig() {
}

// TestHttpClient :
func TestEsConfiguration(t *testing.T) {
	suite.Run(t, new(EsAPISuite))
}
