package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
)

// MySQLConns 涉及到多集群的, 无须全局长连接
var MySQLConns = map[string]*MySQLConn{}

type MySQLConn struct {
	Client    *gorm.DB
	IsMigrate bool
}

//func NewMySQLConns() error {
//	for region, info := range config.Conf.Backend.MysqlConfig.Clusters {
//		MySQLConns[region] = NewMysqlConn(info)
//		if info.IsMigration {
//			if db := MySQLConns[region].Client.AutoMigrate(nil); db.Error != nil {
//				return fmt.Errorf("mysql migrate error: %v", db.Error)
//			}
//		}
//	}
//
//	return nil
//}

func NewMysqlConn(conf *config.MysqlInfo) *MySQLConn {
	var (
		err error
	)
	client, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		conf.User, conf.Password, conf.Address, conf.Name, "Asia%2FShanghai"))
	if err != nil {
		logrus.Errorf("mysql connect error: %v", err)
	}
	return &MySQLConn{
		Client:    client,
		IsMigrate: conf.IsMigration,
	}
}
