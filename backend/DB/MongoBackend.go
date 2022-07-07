package DB

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"goOrigin/backend"
	"goOrigin/config"
)

type MonBackend struct {
	*mgo.Session
}

func (b *MonBackend) NewConn(config config.Config) backend.Connection {
	mongConf := config.Backend.MongoBackendConfig
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", mongConf.User, mongConf.Password, mongConf.Address, mongConf.Port, mongConf.DB)
	session, err := mgo.Dial(url)
	if err != nil {
		logrus.Errorf("init mongo db fail: %s", err)
	}
	err = session.Ping()
	if err != nil {
		logrus.Errorf("init mongo db fail: %s", err)
	}
	return &MonBackend{Session: session}
}

func (b *MonBackend) Create() (interface{}, error) {
	panic("implement me")
}

func (b *MonBackend) Update() (interface{}, error) {
	panic("implement me")
}

func (b *MonBackend) Delete() (interface{}, error) {
	panic("implement me")
}

func (b *MonBackend) Close() error {
	panic("implement me")
}

func (b *MonBackend) GetCollection(name string) *mgo.Collection {
	return b.DB("").C(name)
}
