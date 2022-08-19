package storage

import (
	"github.com/samuel/go-zookeeper/zk"
	"goOrigin/config"
	"time"
)

var ZKFlags = map[string]int32{
	"Permanent": 0,
	"Ephemeral": zk.FlagEphemeral,
}

const prefix = "/ian/"

type ZKConnection struct {
	*zk.Conn
}

func NewZkConn() error {
	conn, _, err := zk.Connect(config.Conf.Backend.ZKConfig.Address, 5*time.Second)
	if err != nil {
		return err
	}
	GlobalZKConn.Conn = conn
	return nil
}

func (c *ZKConnection) Add(path string, data []byte, flags int32) (string, error) {
	return c.Conn.Create(prefix+path, data, flags, nil)
}

func (c *ZKConnection) Get(path string, data []byte, flags int32) (string, error) {
	return c.Conn.Create(prefix+path, data, flags, nil)
}
