package storage

var Mongo *MongoConn
var ZKConn *ZKConnection
var MySQL *MySQLConn
var RedisCon *RedisConn

type Conn interface {
	Close() error
	InitData(mode string) error
}

func InitMongo() error {
	Mongo = NewMongoConn()
	return nil
}

func InitMySQL() error {
	MySQL = InitMySQConn()
	return nil
}

func InitZk() error {
	return NewZkConn()
}

func InitRedis() error {
	RedisCon = NewRedisConn()
	return nil
}
