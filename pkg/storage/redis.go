package storage

import (
	goredis "github.com/go-redis/redis"
	"goOrigin/config"
)

type RedisConn struct {
	*goredis.Client
}

func (r *RedisConn) Close() error {
	panic("implement me")
}

func (r *RedisConn) Exec() ([]byte, error) {
	panic("implement me")
}

func NewRedisConn() *RedisConn {
	return &RedisConn{
		goredis.NewClient(&goredis.Options{
			Network:  "tcp",
			Addr:     config.Conf.Backend.RedisConfig.Addr,
			Password: config.Conf.Backend.RedisConfig.Auth,
		}),
	}
}
