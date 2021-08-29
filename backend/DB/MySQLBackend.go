package DB

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goOrigin/config"
	"goOrigin/logging"
	"goOrigin/run"
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
	var logger = logging.GetStdLogger()

	if conf == nil {
		conf = run.Conf
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
		logger.Fatalf("init mysql client error: %s", err)
		return nil, err
	}

	return &MySQLBackend{
		Client: db,
	}, err
}
