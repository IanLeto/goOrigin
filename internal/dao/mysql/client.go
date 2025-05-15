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
	conf := config.ConfV2
	for region, info := range conf.Env {
		GlobalMySQLConns[region] = NewMysqlV2Conn(info.MysqlSQLConfig)

		if _, ok := GlobalMySQLConns[region]; ok && GlobalMySQLConns[region] != nil {
			if GlobalMySQLConns[region].IsMigrate {
				err := GlobalMySQLConns[region].Migrate()
				if err != nil {
					logger.Sugar().Warnf("环境 %s mysql migrate error: %v", region, err)
				}
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
	logger.Sugar().Warnf("开始mirgrate:")
	err := m.Client.AutoMigrate(dao.TRecord{}, dao.EcampTransTypeTb{}, dao.EcampServiceCodeTb{}, dao.TAchievementRecord{}, dao.TNode{})
	return err
}

func NewMysqlV2Conn(conf config.MySQLConfig) *MySQLConn {
	client, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Address, conf.DBName)), &gorm.Config{
		DisableAutomaticPing: true, // 可选，禁用自动ping

	})
	client.Debug().Session(&gorm.Session{
		NewDB:                    true,
		Initialized:              false,
		SkipHooks:                false,
		SkipDefaultTransaction:   false,
		DisableNestedTransaction: false,
		AllowGlobalUpdate:        false,
		FullSaveAssociations:     false,
		PropagateUnscoped:        false,
		QueryFields:              false,
	})
	if err != nil {
		logger.Sugar().Warnf("❌ MySQL 连接失败: %v", err)
		return nil // ✅ 返回错误，而不是 `nil`
	}

	return &MySQLConn{
		Client:    client,
		IsMigrate: conf.IsMigration,
	}
}
