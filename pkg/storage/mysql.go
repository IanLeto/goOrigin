package storage

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
)

type MySQLConn struct {
	*gorm.DB
}

func InitMySQConn() *MySQLConn {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		config.Conf.Backend.MySqlConfig.User, config.Conf.Backend.MySqlConfig.Password,
		config.Conf.Backend.MySqlConfig.Address, config.Conf.Backend.MySqlConfig.Name, "Asia%2FShanghai"))
	if err != nil {
		logrus.Errorf("init db conn fail %s", err)
	}
	return &MySQLConn{db}

}
