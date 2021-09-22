package storage

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"goOrigin/pkg/utils"
)

var tables = []string{"ian", "todo"}
var tablesIndex = map[string]mgo.Index{
	"ian": mgo.Index{
		Key:              []string{"id"},
		Unique:           true,
		DropDups:         false,
		Background:       false,
		Sparse:           false,
		PartialFilter:    nil,
		ExpireAfter:      0,
		Name:             "",
		Min:              0,
		Max:              0,
		Minf:             0,
		Maxf:             0,
		BucketSize:       0,
		Bits:             0,
		DefaultLanguage:  "",
		LanguageOverride: "",
		Weights:          nil,
		Collation:        nil,
	},
}

type MongoConn struct {
	*mgo.Session
	*mgo.Database
}

func NewMongoConn() *MongoConn {
	mongConf := config.Conf.Backend.MongoBackendConfig
	url := fmt.Sprintf("mongodb://%s:%s/%s", mongConf.Address, mongConf.Port, mongConf.DB)
	session, err := mgo.Dial(url)
	if err != nil {
		logrus.Errorf("init mongo db fail: %s", err)
	}
	err = session.Ping()
	if err != nil {
		logrus.Errorf("init mongo db fail: %s", err)
	}
	return &MongoConn{
		Session:  session,
		Database: session.DB(config.Conf.Backend.MongoBackendConfig.DB),
	}
}

func initSchema(mode string) error {
	names, err := Mongo.DB("ian").CollectionNames()
	if err != nil {
		return err
	}
	var init = func() error {
		for _, t := range tables {
			if !utils.IncludeString(names, t) {
				if err := Mongo.DB("ian").C(t).Create(&mgo.CollectionInfo{}); err != nil {
					return err
				}
			}
		}
		return nil
	}
	switch mode {
	default:
		if err := init(); err != nil {
			return err
		}
		Mongo.DB("ian").C("ian")
	}
	return nil
}

func CreateIndex() {
	for _, table := range tables {
		indexes, err := Mongo.DB("ian").C(table).Indexes()
		if err != nil {
			logrus.Errorf("create index fail %s", err)
		}
		if len(indexes) != 0 {
			continue
		} else {

		}
	}

}
