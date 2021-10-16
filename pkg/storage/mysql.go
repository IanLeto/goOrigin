package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
)

type MySQLConn struct {
	*gorm.DB
}

func InitMySQConn() *MySQLConn {
	db, err := gorm.Open("TCP", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		config.Conf.Backend.MySqlBackendConfig.User, config.Conf.Backend.MySqlBackendConfig.Password,
		config.Conf.Backend.MySqlBackendConfig.Address, config.Conf.Backend.MySqlBackendConfig.Name, "Asia%2FShanghai"))
	if err != nil {
		logrus.Errorf("init db conn fail %s", err)
	}
	return &MySQLConn{db}

}
