package clients

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
)

// MySQLConns 涉及到多集群的, 无须全局长连接
var MySQLConns = map[string]*MySQLConn{}

type MySQLConn struct {
	Client *gorm.DB
}

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
		Client: client,
	}
}
