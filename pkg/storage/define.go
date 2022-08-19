package storage

import "context"

var GlobalMongo *MongoConn
var GlobalZKConn *ZKConnection
var GlobalMySQL *MySQLConn
var GlobalRedisCon *RedisConn

type Conn interface {
	Close() error
	InitData(mode string) error
}

func InitMongo() error {
	// 这里应该加上根ctx
	GlobalMongo = NewMongoConn(context.TODO(), nil)
	return nil
}

func InitMySQL() error {
	GlobalMySQL = InitMySQConn()
	return nil
}

func InitZk() error {
	return NewZkConn()
}

func InitRedis() error {
	GlobalRedisCon = NewRedisConn()
	return nil
}
