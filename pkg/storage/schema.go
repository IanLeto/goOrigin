package storage

//var ianSchema = &mgo.CollectionInfo{
//	DisableIdIndex:   false,
//	ForceIdIndex:     false,
//	Capped:           false,
//	MaxBytes:         0,
//	MaxDocs:          0,
//	Validator:        nil,
//	ValidationLevel:  "",
//	ValidationAction: "",
//	StorageEngine:    nil,
//	Collation:        nil,
//}
//
//var IanInitSchema = func(m *MongoConn) error {
//	c := m.DB(config.Conf.Backend.MongoBackendConfig.DB).C("ian")
//	if err := c.Create(ianSchema); err != nil {
//		return err
//	}
//	if err := c.EnsureIndex(mgo.Index{
//		Key:    []string{"id"},
//		Unique: true,
//	}); err != nil {
//		return err
//	}
//	res, err := ioutil.ReadFile(utils.GetFilePath("internal/model/ian.json"))
//	if err != nil {
//		logrus.Errorf("import data fail %s", err)
//		return err
//	}
//	doc, err := utils.ConvBson(string(res))
//	if err != nil {
//		logrus.Errorf("conv data fail %s", err)
//		return err
//	}
//	err = Mongo.DB("ian").C("ian").Insert(doc)
//	if err != nil {
//		logrus.Errorf("insert data fail %s", err)
//		return err
//	}
//	return nil
//}
//
//var ianIndexCheck = func(m *MongoConn, collection string) bool {
//	c := m.DB(config.Conf.Backend.MongoBackendConfig.DB).C(collection)
//	index, err := c.Indexes()
//	if err != nil {
//		logrus.Errorf("获取index 失败 %s", err)
//	}
//	if len(index) == 0 {
//		return false
//	}
//	for _, i := range index {
//		logrus.Debug(json.Marshal(i))
//	}
//	return true
//}
//
//func InitSchema(m *MongoConn, collectionName string) error {
//	c := m.DB(config.Conf.Backend.MongoBackendConfig.DB).C(collectionName)
//	if err := c.Create(ianSchema); err != nil {
//		if c.Name == "" {
//			return err
//		}
//	}
//	if err := c.EnsureIndex(mgo.Index{
//		Key:    []string{"id"},
//		Unique: true,
//	}); err != nil {
//		return err
//	}
//	res, err := ioutil.ReadFile(utils.GetFilePath("sql/mongoData/testIndex2.json"))
//	if err != nil {
//		logrus.Errorf("import data fail %s", err)
//		return err
//	}
//	doc, err := utils.ConvBson(string(res))
//	if err != nil {
//		logrus.Errorf("conv data fail %s", err)
//		return err
//	}
//	err = Mongo.DB("ian").C(collectionName).Insert(doc)
//	if err != nil {
//		logrus.Errorf("insert data fail %s", err)
//		return err
//	}
//	return nil
//}
