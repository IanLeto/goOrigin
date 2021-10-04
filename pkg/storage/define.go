package storage

var Mongo *MongoConn
var ZKConn *ZKConnection

type Conn interface {
	Close() error
	InitData(mode string) error
}

func InitMongo() error {
	Mongo = NewMongoConn()
	return nil
}

func InitZk() error {
	return NewZkConn()
}
