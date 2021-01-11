package backend

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goOrigin/conf"
	"goOrigin/logging"
)

type MySQLBackend struct {
	Client
}

func NewMySQLBackend(config string) (*gorm.DB, error) {
	var logger = logging.GetStdLogger()
	c := conf.Conf
	if config == "" {
		config = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
			c.Backend.MySqlBackendConfig.User,
			c.Backend.MySqlBackendConfig.Password,
			c.Backend.MySqlBackendConfig.Address,
			"localDB",
			true,
			"Local")
		//"root:root@tcp(localhost:3308)/gormloc2?parseTime=true")
	}

	db, err := gorm.Open("mysql", config)
	if err != nil {
		logger.Fatalf("init mysql db error: %s", err)
		return nil, err
	}

	return db, err
}
