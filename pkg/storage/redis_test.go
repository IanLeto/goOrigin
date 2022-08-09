package storage_test

import (
	"goOrigin/config"
	"goOrigin/pkg/storage"
	"goOrigin/testsuit"
	"testing"

	"github.com/stretchr/testify/suite"
)

// RedisSuite :
type RedisSuite struct {
	suite.Suite
	conf *config.Config
}

func (s *RedisSuite) SetupTest() {
	testsuit.InitTestConfig(config.Config{
		Name:    "",
		Port:    "",
		RunMode: "",
		Backend: &config.BackendConfig{
			MySqlBackendConfig: nil,
			RedisConfig: &config.RedisConfig{
				DB:         0,
				Addr:       "124.222.48.125:6379",
				IsSentinel: false,
				Auth:       "i012nsl9",
			},
			ZKConfig: nil,
		},
		Client:     nil,
		Components: nil,
	})
	s.NoError(storage.InitRedis())
}

// TestMarshal :
func (s *RedisSuite) TestConfig() {
	s.NoError(storage.RedisCon.Ping().Err())
	pong, err := storage.RedisCon.Ping().Result()
	s.Equal("PONG", pong)
	s.NoError(err)
}

// TestHttpClient :
func TestRedisConfiguration(t *testing.T) {
	suite.Run(t, new(RedisSuite))
}
