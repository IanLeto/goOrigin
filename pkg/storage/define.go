package storage

var Mongo *MongoConn
var ZKConn *ZKConnection

func InitMongo() error {
	Mongo = NewMongoConn()
	return nil
}

func InitZk() error {
	return NewZkConn()
}
