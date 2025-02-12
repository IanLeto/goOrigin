package mysql

import (
	"fmt"
	"goOrigin/config"
	"goOrigin/internal/model/dao"
	logger2 "goOrigin/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var logger, _ = logger2.InitZap()

// GlobalMySQLConns 涉及到多集群的, 无须全局长连接
var GlobalMySQLConns = map[string]*MySQLConn{}

func NewMySQLConns() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("⚠️ 数据库迁移失败 : %v\n", r)
		}
	}()
	conf := config.ConfV2
	for region, info := range conf.Env {
		GlobalMySQLConns[region] = NewMysqlV2Conn(info.MysqlSQLConfig)
		if GlobalMySQLConns[region].IsMigrate {
			err := GlobalMySQLConns[region].Migrate()
			if err != nil {
				logger.Sugar().Warnf("环境 %s mysql migrate error: %v", region, err)
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
	err := m.Client.AutoMigrate(dao.TRecord{}, dao.TTransInfo{})
	return err
}

func NewMysqlV2Conn(conf config.MySQLConfig) *MySQLConn {
	var (
		err    error
		client *gorm.DB
	)

	client, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		conf.User, conf.Password, conf.Address, conf.DBName, "Asia%2FShanghai")), &gorm.Config{})

	if err != nil {
		logger.Sugar().Warnf("mysql connect error: %v", err)
	}
	// 新建GORM配置对象，设置日志级别为Info
	return &MySQLConn{
		Client:    client,
		IsMigrate: conf.IsMigration,
	}
}
