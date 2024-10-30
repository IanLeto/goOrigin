package mysql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"goOrigin/internal/model/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// GlobalMySQLConns 涉及到多集群的, 无须全局长连接
var GlobalMySQLConns = map[string]*MySQLConn{}

func NewMySQLConns() error {
	conf := config.ConfV2
	for region, info := range conf.Env {

		GlobalMySQLConns[region] = NewMysqlV2Conn(info.MysqlSQLConfig)
		if GlobalMySQLConns[region].IsMigrate {
			err := GlobalMySQLConns[region].Migrate()
			if err != nil {
				logrus.Errorf("mysql migrate error: %v", err)
			}
		}

	}
	return nil
}

type MySQLConn struct {
	Client    *gorm.DB
	IsMigrate bool
}

func (m *MySQLConn) Migrate() error {
	return m.Client.AutoMigrate(dao.TRecord{})
}

func NewMysqlV2Conn(conf config.MySQLConfig) *MySQLConn {
	var (
		err    error
		client *gorm.DB
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出到标准输出）
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 禁用彩色打印
		},
	)

	client, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		conf.User, conf.Password, conf.Address, conf.DBName, "Asia%2FShanghai")), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		logrus.Errorf("mysql connect error: %v", err)
	}
	// 新建GORM配置对象，设置日志级别为Info
	return &MySQLConn{
		Client:    client,
		IsMigrate: conf.IsMigration,
	}
}

func NewMysqlConn(conf *config.MysqlInfo) *MySQLConn {
	var (
		err error
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出到标准输出）
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 禁用彩色打印
		},
	)

	client, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		conf.User, conf.Password, conf.Address, conf.Name, "Asia%2FShanghai")), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		logrus.Errorf("mysql connect error: %v", err)
	}
	// 新建GORM配置对象，设置日志级别为Info

	return &MySQLConn{
		Client:    client,
		IsMigrate: conf.IsMigration,
	}
}
