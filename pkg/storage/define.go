package storage

import "context"

var Mongo *MongoConn
var ZKConn *ZKConnection
var MySQL *MySQLConn
var RedisCon *RedisConn

type Conn interface {
	Close() error
	InitData(mode string) error
}

func InitMongo() error {
	// 这里应该加上根ctx
	Mongo = NewMongoConn(context.TODO(), nil)
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
