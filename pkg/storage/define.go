package storage

var Mongo *MongoConn

func InitMongo() {
	Mongo = NewMongoConn()
}
