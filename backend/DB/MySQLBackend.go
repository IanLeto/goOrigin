package DB

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goOrigin/config"
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

func NewMySQLBackend(conf *config.Config) (*MySQLBackend, error) {

	if conf == nil {
		conf = config.Conf
	}
	mysqlConf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		conf.Backend.MySqlBackendConfig.User,
		conf.Backend.MySqlBackendConfig.Password,
		conf.Backend.MySqlBackendConfig.Address,
		"localDB",
		true,
		"Local")

	db, err := gorm.Open("mysql", mysqlConf)
	if err != nil {
		return nil, err
	}

	return &MySQLBackend{
		Client: db,
	}, err
}
