package backend

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goOrigin/config"
	"goOrigin/logging"
)

type MySQLBackend struct {
	Client *gorm.DB
}

func (b *MySQLBackend) ShowType() string {
	return "MySQL"
}

func (b *MySQLBackend) Close() error {
	return b.Client.Close()
}

func NewMySQLBackend(conf string) (*MySQLBackend, error) {
	var logger = logging.GetStdLogger()
	c := config.GlobalConfig
	if conf == "" {
		conf = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
			c.Backend.MySqlBackendConfig.User,
			c.Backend.MySqlBackendConfig.Password,
			c.Backend.MySqlBackendConfig.Address,
			"localDB",
			true,
			"Local")
	}

	db, err := gorm.Open("mysql", conf)
	if err != nil {
		logger.Fatalf("init mysql client error: %s", err)
		return nil, err
	}

	return &MySQLBackend{
		Client: db,
	}, err
}
