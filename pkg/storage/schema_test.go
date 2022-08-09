package storage_test

import (
	"goOrigin/config"
	"goOrigin/pkg/storage"
	"goOrigin/testsuit"
	"testing"

	"github.com/stretchr/testify/suite"
)

// MongoSuite :
type MongoSuite struct {
	suite.Suite
	conf *config.Config
}

func (s *MongoSuite) SetupTest() {
	testsuit.InitTestConfig(config.Config{
		Name:    "",
		Port:    "",
		RunMode: "",
		Backend: &config.BackendConfig{
			MySqlBackendConfig: nil,
			MongoBackendConfig: &config.MongoBackendConfig{
				Address:  "localhost:27017",
				Port:     "",
				Password: "",
				User:     "",
				DB:       "ian",
			},
			ZKConfig: nil,
		},
		Client:     nil,
		Components: nil,
	})
	s.NoError(storage.InitMongo())
}

type testIndex struct {
}

// TestMarshal :
func (s *MongoSuite) TestConfig() {
}

// TestHttpClient :
func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(MongoSuite))
}
