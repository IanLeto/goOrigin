package storage

import (
	"github.com/sirupsen/logrus"
)

var Mongo *MongoConn
var ZKConn *ZKConnection

func InitMongo() error {
	Mongo = NewMongoConn()
	err := initSchema("init")
	if err != nil {
		logrus.Error(err)
	}
	return nil
}

func InitZk() error {
	return NewZkConn()
}
