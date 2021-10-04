package storage

import (
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"goOrigin/pkg/utils"
	"io/ioutil"
)

var ianSchema = &mgo.CollectionInfo{
	DisableIdIndex:   false,
	ForceIdIndex:     false,
	Capped:           false,
	MaxBytes:         0,
	MaxDocs:          0,
	Validator:        nil,
	ValidationLevel:  "",
	ValidationAction: "",
	StorageEngine:    nil,
	Collation:        nil,
}

var ianInitSchema = func(m *MongoConn) error {
	c := m.DB(config.Conf.Backend.MongoBackendConfig.DB).C("ian")
	if err := c.Create(ianSchema); err != nil {
		return err
	}
	if err := c.EnsureIndex(mgo.Index{
		Key:    []string{"id"},
		Unique: true,
	}); err != nil {
		return err
	}
	res, err := ioutil.ReadFile(utils.GetFilePath("internal/model/ian.json"))
	if err != nil {
		logrus.Errorf("import data fail %s", err)
		return err
	}
	doc, err := utils.ConvBson(string(res))
	if err != nil {
		logrus.Errorf("conv data fail %s", err)
		return err
	}
	err = Mongo.DB("ian").C("ian").Insert(doc)
	if err != nil {
		logrus.Errorf("insert data fail %s", err)
		return err
	}
	return nil
}
