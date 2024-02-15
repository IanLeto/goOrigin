package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/config"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
	"goOrigin/internal/model/repository"
)

// import (
//
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/olivere/elastic/v7"
//	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/client_golang/prometheus/push"
//	"github.com/sirupsen/logrus"
//	"github.com/spf13/cast"
//	"go.mongodb.org/mongo-driver/bson"
//	"goOrigin/API/V1"
//	"goOrigin/config"
//	"goOrigin/internal/dao/mysql"
//
//	"goOrigin/pkg/clients"
//	logger2 "goOrigin/pkg/logger"
//	"goOrigin/pkg/storage"
//	"time"
//
// )
func CreateRecord(ctx *gin.Context, record *entity.Record) (id uint, err error) {
	var (
		tRecord = &dao.TRecord{}
	)
	tRecord = repository.ToRecordDAO(record)
	db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[record.Region])
	res, _, err := mysql.Create(db.Client, tRecord)
	if err != nil {
		logrus.Errorf("create record failed %s: %s", err, res)
		goto ERR
	}
	return tRecord.ID, err
ERR:
	return 0, err

}

func UpdateRecord(ctx *gin.Context, record *entity.Record) (id uint, err error) {
	var (
		tRecord = &dao.TRecord{}
	)
	tRecord = repository.ToRecordDAO(record)
	db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[record.Region])
	res, _, err := mysql.Create(db.Client, tRecord)
	if err != nil {
		logrus.Errorf("create record failed %s: %s", err, res)
		goto ERR
	}
	return tRecord.ID, err
ERR:
	return 0, err

}

func DeleteRecord(ctx *gin.Context, record *entity.Record) (id uint, err error) {
	var (
		tRecord = &dao.TRecord{}
	)
	tRecord = repository.ToRecordDAO(record)
	db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[record.Region])
	res, _, err := mysql.Create(db.Client, tRecord)
	if err != nil {
		logrus.Errorf("create record failed %s: %s", err, res)
		goto ERR
	}
	return tRecord.ID, err
ERR:
	return 0, err

}

//var weight = prometheus.NewGaugeVec(prometheus.GaugeOpts{
//	Name: "ianRecord",
//	Help: "record ian",
//}, []string{"BF", "LUN", "DIN", "EX"})
//
//func newPusher(info prometheus.Gauge) *push.Pusher {
//	reg := prometheus.NewRegistry()
//	reg.MustRegister(info)
//	return push.New(config.Conf.Backend.PromConfig.Push, config.Conf.Backend.PromConfig.Group).Gatherer(reg)
//}
//
//func CreateIanRecord(c *gin.Context, req V1.CreateIanRequestInfo) (id interface{}, err error) {
//	var (
//		client *elastic.Client
//		res    *elastic.IndexResponse
//		logger = logger2.NewLogger()
//	)
//	client, err = clients.NewESClient()
//	defer func() { client.CloseIndex("ian") }()
//	if err != nil {
//		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
//		return nil, err
//	}
//	ian := entity.NewIan(req)
//	res, err = client.Index().Index("ian").BodyJson(ian).Do(c)
//	if err != nil {
//		logger.Error(fmt.Sprintf("创建 ianrecord:%s 失败 %s", ian.ToString(), err))
//		return nil, err
//	}
//	if req.Body.Weight == 0 {
//		return res, err
//	}
//	// 不另启dao了 写入prometheus
//	info := weight.WithLabelValues(req.Body.BF, req.Body.LUN, req.Body.DIN, req.Body.EXTRA)
//	info.Set(cast.ToFloat64(req.Body.Weight))
//
//	//res, err := storage.GlobalMongo.DB.Collection("ian").InsertOne(context.TODO(), &ian)
//	//if err != nil {
//	//	logrus.Errorf("创建日常数据失败")
//	//	goto ERR
//	//}
//	if err := newPusher(info).Push(); err != nil {
//		logrus.Errorf("push prom failed %s", err)
//		goto ERR
//	}
//	return nil, err
//	//return res.InsertedID.(primitive.ObjectID).Hex(), nil
//ERR:
//	return "", nil
//}
//
//func DeleteIanRecord(c *gin.Context, id string) (delCount int64, err error) {
//
//	res, err := storage.GlobalMongo.DB.Collection("ian").DeleteMany(context.TODO(), bson.M{"name": id})
//	if err != nil {
//		logrus.Errorf("删除日常数据失败 %s", err)
//		goto ERR
//	}
//	return res.DeletedCount, nil
//ERR:
//	return 0, nil
//
//}
//
//func UpdateIanRecord(c *gin.Context, req V1.CreateIanRequestInfo) (id interface{}, err error) {
//	var (
//		ian = entity.NewIan(req)
//	)
//	info := weight.WithLabelValues(req.Body.BF, req.Body.LUN, req.Body.DIN, req.Body.EXTRA)
//	info.Set(cast.ToFloat64(req.Body.Weight))
//	res := storage.GlobalMongo.DB.Collection("ian").FindOneAndReplace(context.TODO(), bson.M{"name": req.Name},
//		&ian)
//
//	if res.Err() != nil {
//		logrus.Errorf("创建日常数据失败 %s", res.Err())
//		goto ERR
//	}
//	if err := newPusher(info).Push(); err != nil {
//		logrus.Errorf("push prom failed %s", err)
//		goto ERR
//	}
//
//	return req.Name, nil
//ERR:
//	return "", res.Err()
//}
//
////func SelectIanRecord(c *gin.Context, req *params.QueryRequest) (response []*params.QueryResponse, err error) {
////	response = make([]*params.QueryResponse, 0)
////
////	filter := bson.M{
////		"name": req.Name,
////	}
////	if req.Name == "" {
////		filter = bson.M{}
////	}
////	curs, err := storage.GlobalMongo.DB.Collection("ian").Find(context.TODO(), filter)
////	if curs.Err() != nil {
////		logrus.Errorf("查询日常数据失败 %s", curs.Err())
////		goto ERR
////	}
////	if err = curs.All(context.TODO(), &response); err != nil {
////		goto ERR
////	}
////	return response, nil
////ERR:
////	return nil, err
////}
//
//func SelectIanRecord(c *gin.Context, req *V1.QueryRequest) (response []*V1.QueryResponse, err error) {
//	var (
//		bq      = elastic.NewBoolQuery()
//		eq      = elastic.NewExistsQuery("name") // 排除无效脚本
//		queries []elastic.Query
//		client  *elastic.Client
//		daoRes  *elastic.SearchResult
//		logger  = logger2.NewLogger()
//	)
//	client, err = clients.NewESClient()
//	defer func() { client.CloseIndex("ian") }()
//	if err != nil {
//		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
//		return nil, err
//	}
//	if req.Name != "" {
//		queries = append(queries, elastic.NewTermQuery("name", req.Name))
//	}
//	bq.Must(queries...)
//	daoRes, err = client.Search().Index("ian").Query(eq).Query(bq).Do(c)
//	if err != nil {
//		logger.Error(fmt.Sprintf("请求es失败 : %s", err))
//		goto ERR
//	}
//	for _, hit := range daoRes.Hits.Hits {
//		var ephemeralIan entity.Ian
//		err = json.Unmarshal(hit.Source, &ephemeralIan)
//		if err != nil {
//			goto ERR
//		}
//		response = append(response, &V1.QueryResponse{
//			Name:       ephemeralIan.Name,
//			CreateTime: time.Unix(ephemeralIan.Time.T, 0).Format(time.RFC3339),
//			UpdateTime: time.Unix(ephemeralIan.Time.I, 0).Format(time.RFC3339),
//			Body: struct {
//				Weight float32 `json:"weight"`
//			}{
//				Weight: ephemeralIan.Body.Weight,
//			},
//			BETre: struct {
//				Core       int `json:"core"`
//				Runner     int `json:"runner"`
//				Support    int `json:"support"`
//				Squat      int `json:"squat"`
//				EasyBurpee int `json:"easy_burpee"`
//				Chair      int `json:"chair"`
//				Stretch    int `json:"stretch"`
//			}{
//				Core:       ephemeralIan.BETre.Core,
//				Runner:     ephemeralIan.BETre.Runner,
//				Support:    ephemeralIan.BETre.Support,
//				Squat:      ephemeralIan.BETre.Squat,
//				EasyBurpee: ephemeralIan.BETre.EasyBurpee,
//				Chair:      ephemeralIan.BETre.Chair,
//				Stretch:    ephemeralIan.BETre.Stretch,
//			},
//			Worker: struct {
//				Vol1 string `json:"vol1"`
//				Vol2 string `json:"vol2"`
//				Vol3 string `json:"vol3"`
//				Vol4 string `json:"vol4"`
//			}{
//				Vol1: ephemeralIan.Worker.Vol1,
//				Vol2: ephemeralIan.Worker.Vol2,
//				Vol3: ephemeralIan.Worker.Vol3,
//				Vol4: ephemeralIan.Worker.Vol4,
//			},
//		})
//
//	}
//
//	return response, nil
//ERR:
//	return nil, err
//}
//
//func AppendIanRecord(c *gin.Context, req *V1.AppendRequestInfo) (*entity.Ian, error) {
//	var (
//		bq      = elastic.NewBoolQuery()
//		queries []elastic.Query
//		client  *elastic.Client
//		daoRes  *elastic.SearchResult
//		ian     entity.Ian
//		err     error
//		logger  = logger2.NewLogger()
//	)
//	client, err = clients.NewESClient()
//	if err != nil {
//		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
//		return nil, err
//	}
//	queries = append(queries, elastic.NewTermQuery("id", req.Id))
//	bq.Must(queries...)
//	daoRes, err = client.Search("ian").Query(bq).Size(1).Do(c)
//	if len(daoRes.Hits.Hits) != 1 {
//		goto ERR
//	}
//	if err := json.Unmarshal(daoRes.Hits.Hits[0].Source, &ian); err != nil {
//		goto ERR
//	}
//	ian.Worker.Vol1 = fmt.Sprintf("%s;%s", ian.Worker.Vol1, req.Worker.Vol1)
//	ian.Worker.Vol2 = fmt.Sprintf("%s;%s", ian.Worker.Vol2, req.Worker.Vol2)
//	ian.Worker.Vol3 = fmt.Sprintf("%s;%s", ian.Worker.Vol3, req.Worker.Vol3)
//	ian.Worker.Vol4 = fmt.Sprintf("%s;%s", ian.Worker.Vol4, req.Worker.Vol4)
//
//	return &ian, err
//
//ERR:
//	return nil, nil
//}
//
////func CreateIanRecordV2(c *gin.Context, req *V1.CreateIanRecordRequest) (*V1.CreateIanRecordResponse, error) {
////	var (
////		err error
////	)
////	tRecord := dao.TRecord{
////		Name:       req.Name,
////		Weight:     req.Weight,
////		BF:         req.BF,
////		LUN:        req.LUN,
////		DIN:        req.DIN,
////		EXTRA:      req.EXTRA,
////		Core:       req.Core,
////		Runner:     req.Runner,
////		Support:    req.Support,
////		Squat:      req.Squat,
////		EasyBurpee: req.EasyBurpee,
////		Chair:      req.Chair,
////		Stretch:    req.Stretch,
////		Vol1:       req.Vol1,
////		Vol2:       req.Vol2,
////		Vol3:       req.Vol3,
////		Vol4:       req.Vol4,
////		Content:    req.Content,
////		Region:     req.Region,
////	}
////	db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[req.Region])
////	res, _, err := mysql.Create(db.Client, &tRecord)
////	if err != nil {
////		logrus.Errorf("create record failed %s: %s", err, res)
////		return &V1.CreateIanRecordResponse{}, err
////	}
////	return &V1.CreateIanRecordResponse{
////		Id: tRecord.ID,
////	}, nil
////
////}
//
//func BatchCreateIanRecordsV2(c *gin.Context, req *V1.BatchCreateIanRecordRequest) (*V1.BatchCreateIanRecordResponse, error) {
//	var (
//		dbs map[string][]*dao.TRecord
//		res = &V1.BatchCreateIanRecordResponse{}
//		err error
//	)
//
//	for _, item := range req.Items {
//		tRecord := dao.TRecord{
//			Name:       item.Name,
//			Weight:     item.Weight,
//			BF:         item.BF,
//			LUN:        item.LUN,
//			DIN:        item.DIN,
//			EXTRA:      item.EXTRA,
//			Core:       item.Core,
//			Runner:     item.Runner,
//			Support:    item.Support,
//			Squat:      item.Squat,
//			EasyBurpee: item.EasyBurpee,
//			Chair:      item.Chair,
//			Stretch:    item.Stretch,
//			Vol1:       item.Vol1,
//			Vol2:       item.Vol2,
//			Vol3:       item.Vol3,
//			Vol4:       item.Vol4,
//			Content:    item.Content,
//			Region:     item.Region,
//		}
//		dbs[item.Region] = append(dbs[item.Region], &tRecord)
//	}
//	for region, records := range dbs {
//		db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region])
//		result, _, err := mysql.BatchCreate(db.Client, records)
//		if err != nil {
//			logrus.Errorf("create record failed %s: %s", err, res)
//			return res, err
//		}
//		res.Items = append(res.Items, result)
//	}
//
//	return res, err
//}
//
////func QueryIanRecordsV2(c *gin.Context, region string, name string, startTime, modifyTime int64, limit int) (*V1.SelectIanRecordResponse, error) {
////	var (
////		err     error
////		records []*dao.TRecord
////		where   = "1 = 1"
////		res     = &V1.SelectIanRecordResponse{
////			Items: []V1.IanRecordInfo{},
////		}
////	)
////	if name != "" {
////		where = fmt.Sprintf("%s and name = '%s'", where, name)
////	}
////	if startTime != 0 {
////		where = fmt.Sprintf("%s and create_time >= %d", where, startTime)
////	}
////	if modifyTime != 0 {
////		where = fmt.Sprintf("%s and update_time >= %d", where, modifyTime)
////	}
////
////	db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region])
////	result, _, err := mysql.GetValueByRaw(db.Client, &records, "t_records", where)
////	if err != nil {
////		logrus.Errorf("create record failed %s: %s", err, result)
////		return nil, err
////	}
////	for _, record := range records {
////		res.Items = append(res.Items, V1.IanRecordInfo{
////			Id:         record.ID,
////			CreateTime: record.CreateTime,
////			ModifyTime: record.ModifyTime,
////			Name:       record.Name,
////			Weight:     record.Weight,
////			BF:         record.BF,
////			LUN:        record.LUN,
////			DIN:        record.DIN,
////			EXTRA:      record.EXTRA,
////			Core:       record.Core,
////			Runner:     record.Runner,
////			Support:    record.Support,
////			Squat:      record.Squat,
////			EasyBurpee: record.EasyBurpee,
////			Chair:      record.Chair,
////			Stretch:    record.Stretch,
////			Vol1:       record.Vol1,
////			Vol2:       record.Vol2,
////			Vol3:       record.Vol3,
////			Vol4:       record.Vol4,
////			Content:    record.Content,
////			Region:     record.Region,
////		})
////	}
////
////	return res, err
////
////}
//
//func UpdateIanRecordsV2(c *gin.Context, req *V1.UpdateIanRecordRequest) (res *V1.UpdateIanRecordResponse, err error) {
//	var (
//		record = dao.TRecord{
//			Name:       req.Info.Name,
//			Weight:     req.Info.Weight,
//			BF:         req.Info.BF,
//			LUN:        req.Info.LUN,
//			DIN:        req.Info.DIN,
//			EXTRA:      req.Info.EXTRA,
//			Core:       req.Info.Core,
//			Runner:     req.Info.Runner,
//			Support:    req.Info.Support,
//			Squat:      req.Info.Squat,
//			EasyBurpee: req.Info.EasyBurpee,
//			Chair:      req.Info.Chair,
//			Stretch:    req.Info.Stretch,
//			Vol1:       req.Info.Vol1,
//			Vol2:       req.Info.Vol2,
//			Vol3:       req.Info.Vol3,
//			Vol4:       req.Info.Vol4,
//			Content:    req.Info.Content,
//			Region:     req.Info.Region,
//		}
//	)
//	db := mysql.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[req.Info.Region])
//	result, _, err := mysql.UpdateValue(db.Client, "", "", &record)
//	if err != nil {
//		logrus.Errorf("create record failed %s: %s", err, result)
//		return
//	}
//	res = &V1.UpdateIanRecordResponse{
//		Item: V1.IanRecordInfo{
//			Name:       record.Name,
//			Weight:     record.Weight,
//			BF:         record.BF,
//			LUN:        record.LUN,
//			DIN:        record.DIN,
//			EXTRA:      record.EXTRA,
//			Core:       record.Core,
//			Runner:     record.Runner,
//			Support:    record.Support,
//			Squat:      record.Squat,
//			EasyBurpee: record.EasyBurpee,
//			Chair:      record.Chair,
//			Stretch:    record.Stretch,
//			Vol1:       record.Vol1,
//			Vol2:       record.Vol2,
//			Vol3:       record.Vol3,
//			Vol4:       record.Vol4,
//			Content:    record.Content,
//			Region:     record.Region,
//		},
//	}
//	return
//
//}
