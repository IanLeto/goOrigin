package clients_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"goOrigin/config"
	"goOrigin/pkg/clients"
	"testing"
)

type EsAPISuite struct {
	suite.Suite
	conn *clients.EsV2Conn
}

func (s *EsAPISuite) SetupTest() {
	var err error
	s.conn = clients.NewEsV2Conn(&config.Config{
		Backend: &config.BackendConfig{
			EsConfig: &config.EsConfig{Address: "http://49.233.61.57:9211"},
		},
	})
	s.NoError(err)
}

// TestMarshal :
func (s *EsAPISuite) TestConfig() {
	res, err := s.conn.Client.Info()
	s.NoError(err)
	var (
		buf   bytes.Buffer
		query map[string]interface{}
	)
	query = map[string]interface{}{}
	err = json.NewEncoder(&buf).Encode(query)
	s.NoError(err)
	//buf.Write([]byte(""))
	res, err = s.conn.Client.Search(
		s.conn.Client.Search.WithIndex("audit1"),
		s.conn.Client.Search.WithBody(&buf),
	)
	s.NoError(err)
	defer func() { _ = res.Body.Close() }()
	fmt.Println(res.String())

}

func (s *EsAPISuite) TestQuery() {

}

// TestHttpClient :
func TestEsConfiguration(t *testing.T) {
	suite.Run(t, new(EsAPISuite))
}
