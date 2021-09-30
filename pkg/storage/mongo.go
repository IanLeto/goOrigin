package storage

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"goOrigin/pkg/utils"
)

var tables = map[string]func(m *MongoConn) error{
	"ian": ianInitSchema,
}

type MongoConn struct {
	*mgo.Session
	*mgo.Database
}

func (m *MongoConn) Close() error {
	m.Session.Close()
	return nil
}

func (m *MongoConn) InitData(mode string) error {
	if mode != "init" {
		return nil
	}
	collections, err := m.DB(config.Conf.Backend.MongoBackendConfig.DB).CollectionNames()
	if err != nil {
		logrus.Errorf("init collections failed %s", err)
		return err
	}
	for name, fn := range tables {
		if !utils.IncludeString(collections, name) {
			if err := fn(m); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewMongoConn() *MongoConn {
	mongConf := config.Conf.Backend.MongoBackendConfig
	url := fmt.Sprintf("mongodb://%s:%s/%s", mongConf.Address, mongConf.Port, mongConf.DB)
	session, err := mgo.Dial(url)
	if err != nil {
		logrus.Errorf("init mongo db fail: %s", err)
	}
	err = session.Ping()
	if err != nil {
		logrus.Errorf("init mongo db fail: %s", err)
	}
	return &MongoConn{
		Session:  session,
		Database: session.DB(config.Conf.Backend.MongoBackendConfig.DB),
	}
}


