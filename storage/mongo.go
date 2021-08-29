package storage

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"goOrigin/run"
)

type MongoConn struct {
	*mgo.Session
}

func NewMongoConn() *MongoConn {
	mongConf := run.Conf.Backend.MongoBackendConfig
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", mongConf.User, mongConf.Password, mongConf.Address, mongConf.Port, mongConf.DB)
	session, err := mgo.Dial(url)
	if err != nil {
		logrus.Errorf("init mongo db fail: %s", err)
	}
	err = session.Ping()
	if err != nil {
		logrus.Errorf("init mongo db fail: %s", err)
	}
	return &MongoConn{Session: session}
}
