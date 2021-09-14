package storage

import "goOrigin/config"

var Mongo *MongoConn

func InitMongo() error {
	Mongo = NewMongoConn()
	if config.Conf.RunMode == "init" {

	}
	return nil
}
