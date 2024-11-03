package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goOrigin/API/V1"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
	"goOrigin/pkg/logger"
)

func GetProject(ctx *gin.Context, region string, info *V1.CreateIanRecordRequestInfo) (uint, error) {
	var (
		tRecord      = &dao.TRecord{}
		recordEntity = &entity.Record{}
		logger2, err = logger.InitZap()
	)

	recordEntity.Name = info.Name
	recordEntity.Weight = info.Weight
	recordEntity.Vol1 = info.Vol1
	recordEntity.Vol2 = info.Vol2
	recordEntity.Vol3 = info.Vol3
	recordEntity.Vol4 = info.Vol4
	recordEntity.Content = info.Content
	recordEntity.Retire = info.Retire
	recordEntity.Cost = info.Cost
	recordEntity.Region = region
	recordEntity.Dev = info.Dev
	recordEntity.Coding = info.Coding

	recordEntity.Social = info.Social

	tRecord = repository.ToRecordDAO(recordEntity)
	db := mysql.GlobalMySQLConns[region]
	res, _, err := mysql.Create(db.Client, tRecord)
	if err != nil {
		logger2.Error(fmt.Sprintf("create record failed %s: %s", err, res))
		goto ERR
	}
	//_, err = es.Create("ian", tRecord)

	return tRecord.ID, err
ERR:
	return 0, err

}
