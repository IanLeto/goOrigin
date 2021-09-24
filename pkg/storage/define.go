package storage

import (
	"github.com/sirupsen/logrus"
	"goOrigin/config"
)

var Mongo *MongoConn
var ZKConn *ZKConnection

func InitMongo() error {
	Mongo = NewMongoConn()
	if config.Conf.RunMode == "init" {
		err := initSchema("init")
		if err != nil {
			logrus.Error(err)
		}
	}
	return nil
}

func InitZk() error {
	return NewZkConn()
}
