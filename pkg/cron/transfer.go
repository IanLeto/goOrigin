package cron

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
	"time"
)

type Transfer struct {
	*entity.Record
	Alias string
}

func (t *Transfer) Exec(ctx context.Context) error {
	var (
		esClient = elastic.GlobalEsConns[config.ConfV2.Base.Region]
		index    = t.Alias
		err      error
	)

	// 将Record结构体转换为JSON格式的字节数组
	body, err := json.Marshal(t.Record)
	if err != nil {
		logrus.Errorf("Failed to marshal record: %v", err)
		return err
	}

	// 调用EsV2Conn的Create方法写入数据到Elasticsearch
	resp, err := esClient.Create(index, body)
	if err != nil {
		logrus.Errorf("Failed to create document in Elasticsearch: %v", err)
		return err
	}

	// 处理Elasticsearch的响应
	logrus.Infof("Document created in Elasticsearch. Response: %s", resp)

	return nil
}

func (t *Transfer) GetName() string {
	//TODO implement me
	panic("implement me")
}

// TransferCornFactory : 最重要的方法，如何注册并启动任务
func TransferCornFactory(ctx context.Context) error {
	var (
		interval int
		//err      error
	)
	transferConf := config.ConfV2.Env[config.ConfV2.Base.Region].CronJobConfig.TransferConfig
	interval = transferConf.Interval

	for {
		select {
		case <-time.NewTicker(time.Duration(interval) * time.Second).C:
			var (
				recordTables = make([]*dao.TRecord, 0)
				err          error
			)
			db := mysql.GlobalMySQLConns[config.ConfV2.Base.Region]
			sql := db.Client.Table("t_records")
			sql.Where("created_at > ?", time.Now().Add(-time.Minute))
			tRecords := sql.Find(&recordTables)
			if tRecords.Error != nil {
				logrus.Errorf("create record failed %s: %s", err, tRecords.Error)
				continue
			}
			for _, recordTable := range recordTables {
				GTM.AddJob(&Transfer{
					Record: repository.ToRecordEntity(recordTable),
					Alias:  transferConf.ES.Alias,
				})
			}

		}
	}

	return nil
}
