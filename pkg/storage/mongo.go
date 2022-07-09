package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"goOrigin/config"
)

type MongoConn struct {
	*mongo.Client
	DB *mongo.Database
}

func (m *MongoConn) Close() error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoConn) InitData(mode string) error {
	//TODO implement me
	panic("implement me")
}

func NewMongoConn(ctx context.Context, conf *config.Config) *MongoConn {
	if conf == nil {
		conf = config.Conf
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/%s", conf.Backend.MongoBackendConfig.Address, conf.Backend.MongoBackendConfig.Port, conf.Backend.MongoBackendConfig.DB)))
	client.Database("goOrigin")
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return &MongoConn{client, client.Database(conf.Backend.MongoBackendConfig.DB)}
}

//type MongoConn struct {
//	*mgo.Session
//	*mgo.Database
//}
//
//func (m *MongoConn) Close() error {
//	m.Session.Close()
//	return nil
//}
//
//func (m *MongoConn) InitData(mode string) error {
//	if mode != "init" {
//		return nil
//	}
//	collections, err := m.DB(config.Conf.Backend.MongoBackendConfig.DB).CollectionNames()
//	if err != nil {
//		logrus.Errorf("init collections failed %s", err)
//		return err
//	}
//	for name, fn := range tables {
//		if !utils.IncludeString(collections, name) {
//			if err := fn(m); err != nil {
//				return err
//			}
//		}
//	}
//	logrus.Debug("初始化mongoData完成")
//	return nil
//}
//
//func NewMongoConn() *MongoConn {
//	mongConf := config.Conf.Backend.MongoBackendConfig
//	url := fmt.Sprintf("mongodb://%s:%s/%s", mongConf.Address, mongConf.Port, mongConf.DB)
//	session, err := mgo.Dial(url)
//	if err != nil {
//		logrus.Errorf("init mongo db fail: %s", err)
//	}
//	err = session.Ping()
//	if err != nil {
//		logrus.Errorf("init mongo db fail: %s", err)
//	}
//	return &MongoConn{
//		Session:  session,
//		Database: session.DB(config.Conf.Backend.MongoBackendConfig.DB),
//	}
//}
