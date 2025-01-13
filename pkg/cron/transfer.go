package cron

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"goOrigin/internal/dao/elastic"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
	"goOrigin/pkg/processor"
	"io"
	"os"
	"time"
)

type Transfer struct {
	*entity.RecordEntity
	Alias string
}

func (t *Transfer) Exec(ctx context.Context) error {
	var (
		//esClient = elastic.GlobalEsConns[config.ConfV2.Base.Region]
		//index    = t.Alias
		err error
	)
	file, err := os.Open("./test.json")
	defer func() { _ = file.Close() }()

	// 创建一个缓冲区读取器,用于读取 JSON 数据
	reader := bufio.NewReader(file)

	// 创建一个字节缓冲区,用于存储读取的数据
	var buffer bytes.Buffer

	// 循环读取 JSON 数据,直到读取完毕或出错
	for {
		// 读取一行数据
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				// 读取完毕,跳出循环
				break
			}
			logrus.Errorf("Failed to read record data: %v", err)
			return err
		}

		// 将读取的数据写入字节缓冲区
		buffer.Write(line)
	}

	// 将字节缓冲区中的数据转换为字节数组
	body := buffer.Bytes()
	logrus.Infof("Read record data: %s", body)
	// 调用 EsV2Conn 的 Create 方法写入数据到 Elasticsearch
	//resp, err := esClient.Create(index, body)

	if err != nil {
		logrus.Errorf("Failed to create document in Elasticsearch: %v", err)
		return err
	}

	// 处理 Elasticsearch 的响应
	//logrus.Infof("Document created in Elasticsearch. Response: %s", resp)

	return nil
}

func (t *Transfer) Exec2(ctx context.Context) error {
	var (
		esClient = elastic.GlobalEsConns[config.ConfV2.Base.Region]
		index    = t.Alias
		err      error
	)

	// 将Record结构体转换为JSON格式的字节数组
	body, err := json.Marshal(t.RecordEntity)
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
	return ""
}

// TransferCornFactory : 最重要的方法，如何注册并启动任务
func TransferCornFactory() error {
	for {
		select {
		case <-time.NewTicker(time.Duration(10) * time.Second).C:

			GTM.AddJob(&Transfer{})

		}
	}
	return nil
}
func TransferCornFactory2() error {
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
			logger.Sugar().Infoln("根据每条数据的特性，开始封装数据流水线")
			func() {
				// 读取配置文件，封装数据流水线
				var pipeline = processor.Pipeline{}
				var node1 processor.Node
				pipeline.Add(node1)

			}()

			for _, recordTable := range recordTables {
				GTM.AddJob(&Transfer{
					RecordEntity: repository.ToRecordEntity(recordTable),
					Alias:        transferConf.ES.Alias,
				})
			}

		}
	}

	return nil
}
