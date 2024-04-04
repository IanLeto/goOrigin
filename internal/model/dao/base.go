package dao

type Meta struct {
	ID         uint  `swaggerignore:"true" gorm:"primary_key" json:"id" binding:"-" `
	CreateTime int64 `swaggerignore:"true" gorm:"autoCreateTime;" json:"created_time" binding:"-"`
	ModifyTime int64 `swaggerignore:"true" gorm:"autoUpdateTime;" json:"modify_time" binding:"-"`
}

type Table interface {
	GetID() uint
}

var migrate = map[string]interface{}{
	"t_record": &TRecord{},
}

func DBMigrate() error {
	//for region, _ := range config.Conf.Backend.MysqlConfig.Clusters {
	//	for _, table := range migrate {
	//		//fmt.Println(mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region]).Client.AutoMigrate(table).Error())
	//
	//	}
	//}

	return nil
}
