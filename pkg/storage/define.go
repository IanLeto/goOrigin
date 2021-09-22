package storage

import "goOrigin/config"

var Mongo *MongoConn
var ZKConn *ZKConnection

func InitMongo() error {
	Mongo = NewMongoConn()
	if config.Conf.RunMode == "init" {

	}
	return nil
}

func InitZk() error {
	return NewZkConn()
}
