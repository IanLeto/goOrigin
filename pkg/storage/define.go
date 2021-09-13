package storage

var Mongo *MongoConn

func InitMongo() error {
	Mongo = NewMongoConn()
	return nil
}
