package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/config"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
	"os"
)

func CreateRecord(ctx *gin.Context, info *V1.CreateIanRecordRequestInfo) (uint, error) {
	var (
		tRecord      = &dao.TRecord{}
		recordEntity = &entity.RecordEntity{}
		region       = ctx.GetString("region")
	)
	recordEntity.Title = info.Title
	recordEntity.MorWeight = info.MorWeight
	recordEntity.NigWeight = info.NigWeight
	recordEntity.Vol1 = info.Vol1
	recordEntity.Vol2 = info.Vol2
	recordEntity.Vol3 = info.Vol3
	recordEntity.Vol4 = info.Vol4
	recordEntity.Content = info.Content
	recordEntity.Cost = info.Cost

	tRecord = repository.ToRecordDAO(recordEntity)
	db := mysql.GlobalMySQLConns[region]
	res, _, err := mysql.Create(db.Client, tRecord)
	if err != nil {
		logger.Error(fmt.Sprintf("create record failed %s: %s", err, res))
		goto ERR
	}

	return tRecord.ID, err
ERR:
	return 0, err

}

func CreateFileRecord(ctx *gin.Context, info *V1.CreateIanRecordRequestInfo) (uint, error) {
	var (
		tRecord      = &dao.TRecord{}
		recordEntity = &entity.RecordEntity{}
		region       = ctx.GetString("region")
		path         string
	)

	recordEntity.Title = info.Title
	recordEntity.MorWeight = info.MorWeight
	recordEntity.NigWeight = info.NigWeight
	recordEntity.Vol1 = info.Vol1
	recordEntity.Vol2 = info.Vol2
	recordEntity.Vol3 = info.Vol3
	recordEntity.Vol4 = info.Vol4
	recordEntity.Content = info.Content
	recordEntity.Cost = info.Cost
	recordEntity.Social = info.Social

	tRecord = repository.ToRecordDAO(recordEntity)
	switch region {
	case "win":
		path = "/home/ian/workdir/goOrigin/records.json"
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	jsonData, err := json.Marshal(tRecord)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to open file: %s", err))
		goto ERR
	}

	// 检查文件是否存在,如果不存在则创建

	defer file.Close()

	// 将JSON数据写入文件末尾
	if _, err := file.Write(jsonData); err != nil {
		logger.Error(fmt.Sprintf("failed to write JSON data to file: %s", err))
		goto ERR
	}

	// 添加换行符,方便下一次写入
	if _, err := file.WriteString("\n"); err != nil {
		logger.Error(fmt.Sprintf("failed to write newline character to file: %s", err))
		goto ERR
	}

	return 0, err

ERR:
	return 0, err
}

func UpdateRecord(ctx *gin.Context, region string, record *entity.RecordEntity) (id uint, err error) {
	var tRecord = &dao.TRecord{}

	// 将 entity 转换为 DAO 结构
	tRecord = repository.ToRecordDAO(record)

	// 连接数据库
	db := mysql.NewMysqlV2Conn(config.ConfV2.Env[region].MysqlSQLConfig)

	// 检查记录是否存在
	existingRecord := &dao.TRecord{}
	if err := db.Client.Table("t_records").Where("id = ?", tRecord.ID).First(existingRecord).Error; err != nil {

		logrus.Errorf("query record failed: %v", err)
		return 0, err
	}

	// 更新记录
	if err := db.Client.Table("t_records").Where("id = ?", tRecord.ID).Updates(tRecord).Error; err != nil {
		logrus.Errorf("update record failed: %v", err)
		return 0, err
	}

	return tRecord.ID, nil
}

func QueryRecords(ctx *gin.Context, region string, name string, startTime, endTime int64, pageSize, page int) ([]*entity.RecordEntity, error) {
	var (
		recordEntities = make([]*dao.TRecord, 0)
		err            error
		res            = make([]*entity.RecordEntity, 0)
	)

	offset := (page - 1) * pageSize

	// 获取数据库连接
	db := mysql.GlobalMySQLConns[region]

	// 开启调试模式，打印 SQL 语句
	sql := db.Client.Debug().Table("t_records")

	// 添加查询条件
	if name != "" {
		sql = sql.Where("name = ?", name)
	}
	if startTime != 0 {
		sql = sql.Where("create_time > ?", startTime)
	}
	if endTime != 0 {
		sql = sql.Where("modify_time < ?", endTime)
	}
	if pageSize == 0 {
		pageSize = 50
	}

	// 分页查询
	tRecords := sql.Order("create_time DESC").Limit(pageSize).Offset(offset).Find(&recordEntities).Order("create_time")
	if tRecords.Error != nil {
		logrus.Errorf("query records failed: %s", tRecords.Error)
		goto ERR
	}

	// 转换数据
	for _, recordEntity := range recordEntities {
		res = append(res, repository.ToRecordEntity(recordEntity))
	}

	return res, nil

ERR:
	return nil, err
}

// DeleteRecord 根据 ID 删除记录
func DeleteRecord(ctx *gin.Context, region string, recordID uint) error {
	if region == "" {
		region = "win"
	}

	db := mysql.GlobalMySQLConns[region]

	// 先检查记录是否存在（可选）
	var tRecord dao.TRecord
	if err := db.Client.First(&tRecord, recordID).Error; err != nil {
		logger.Error(fmt.Sprintf("record not found: %s", err))
		return err
	}

	// 执行删除
	if err := db.Client.Delete(&tRecord).Error; err != nil {
		logger.Error(fmt.Sprintf("delete record failed: %s", err))
		return err
	}

	logger.Info(fmt.Sprintf("record deleted successfully: ID = %d", recordID))
	return nil
}
