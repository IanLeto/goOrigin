package storage

import (
	"github.com/sirupsen/logrus"
	"goOrigin/config"
)

var Mongo *MongoConn
var ZKConn *ZKConnection

type Conn interface {
	Close() error
	InitData(mode string) error
}

func InitMongo() error {
	Mongo = NewMongoConn()
	err := initSchema(config.Conf.RunMode)
	if err != nil {
		logrus.Error(err)
	}
	return nil
}

func InitZk() error {
	return NewZkConn()
}
