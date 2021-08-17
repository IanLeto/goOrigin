package DB

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
)

type MonBackend struct {
	*mgo.Session
}

func NewMonBackend(conf config.MongoBackendConfig) *MonBackend {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/", conf.User, conf.Password, conf.Address, conf.Port)
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
