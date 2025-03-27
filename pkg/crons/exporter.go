package crons

import (
	"context"
	"github.com/robfig/cron/v3"
)

func CronGetMetricFromMySQL(ctx context.Context) {
	panic(1)
	//region, _ := conv.String(ctx.Value("region"))
	//var (
	//	db              = mysql.GlobalMySQLConns[region]
	//	recordEntites   []entity.RecordEntity
	//	tRecordEntities []dao.TRecord
	//	pageSize        = 1000 // 每页查询的记录数
	//	pageNumber      = 1    // 当前页码
	//)
	//for {
	//	// 使用 GORM 进行分页查询
	//	ephtRecordEntities := make([]dao.TRecord, 0)
	//	if err := db.Client.Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&ephtRecordEntities).Error; err != nil {
	//		//log.Printf("Failed to fetch records from MySQL: %v", err)
	//		return
	//	}
	//
	//	// 如果当前页没有数据,说明已经查询完所有记录,退出循环
	//	if len(ephtRecordEntities) == 0 {
	//		break
	//	}
	//
	//	// 遍历当前页的查询结果,将 dao.TRecord 转换为 entity.RecordEntity
	//	for _, tRecord := range ephtRecordEntities {
	//		recordEntity := repository.ToRecordEntity(&tRecord)
	//		tRecordEntities = append(ephtRecordEntities, *recordEntity)
	//	}
	//
	//	// 清空 ephtRecordEntities,准备查询下一页
	//	ephtRecordEntities = ephtRecordEntities[:0]
	//
	//	// 页码加1,查询下一页
	//	pageNumber++
	//}
}

func NewMetricExporter() {
	c := cron.New()
	spec := ""
	c.AddFunc(spec, func() {

	})
	c.Start()
}
