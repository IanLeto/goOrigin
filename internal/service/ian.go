package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"goOrigin/API/V1"
	"goOrigin/config"
	"goOrigin/internal/dao/mysql"
	"goOrigin/internal/model"
	"goOrigin/pkg/clients"
	logger2 "goOrigin/pkg/logger"
	"goOrigin/pkg/storage"
	"time"
)

var weight = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "ianRecord",
	Help: "record ian",
}, []string{"BF", "LUN", "DIN", "EX"})

func newPusher(info prometheus.Gauge) *push.Pusher {
	reg := prometheus.NewRegistry()
	reg.MustRegister(info)
	return push.New(config.Conf.Backend.PromConfig.Push, config.Conf.Backend.PromConfig.Group).Gatherer(reg)
}

func CreateIanRecord(c *gin.Context, req V1.CreateIanRequestInfo) (id interface{}, err error) {
	var (
		client *elastic.Client
		res    *elastic.IndexResponse
		logger = logger2.NewLogger()
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex("ian") }()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return nil, err
	}
	ian := model.NewIan(req)
	res, err = client.Index().Index("ian").BodyJson(ian).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("创建 ianrecord:%s 失败 %s", ian.ToString(), err))
		return nil, err
	}
	if req.Body.Weight == 0 {
		return res, err
	}
	// 不另启dao了 写入prometheus
	info := weight.WithLabelValues(req.Body.BF, req.Body.LUN, req.Body.DIN, req.Body.EXTRA)
	info.Set(cast.ToFloat64(req.Body.Weight))

	//res, err := storage.GlobalMongo.DB.Collection("ian").InsertOne(context.TODO(), &ian)
	//if err != nil {
	//	logrus.Errorf("创建日常数据失败")
	//	goto ERR
	//}
	if err := newPusher(info).Push(); err != nil {
		logrus.Errorf("push prom failed %s", err)
		goto ERR
	}
	return nil, err
	//return res.InsertedID.(primitive.ObjectID).Hex(), nil
ERR:
	return "", nil
}

func DeleteIanRecord(c *gin.Context, id string) (delCount int64, err error) {

	res, err := storage.GlobalMongo.DB.Collection("ian").DeleteMany(context.TODO(), bson.M{"name": id})
	if err != nil {
		logrus.Errorf("删除日常数据失败 %s", err)
		goto ERR
	}
	return res.DeletedCount, nil
ERR:
	return 0, nil

}

func UpdateIanRecord(c *gin.Context, req V1.CreateIanRequestInfo) (id interface{}, err error) {
	var (
		ian = model.NewIan(req)
	)
	info := weight.WithLabelValues(req.Body.BF, req.Body.LUN, req.Body.DIN, req.Body.EXTRA)
	info.Set(cast.ToFloat64(req.Body.Weight))
	res := storage.GlobalMongo.DB.Collection("ian").FindOneAndReplace(context.TODO(), bson.M{"name": req.Name},
		&ian)

	if res.Err() != nil {
		logrus.Errorf("创建日常数据失败 %s", res.Err())
		goto ERR
	}
	if err := newPusher(info).Push(); err != nil {
		logrus.Errorf("push prom failed %s", err)
		goto ERR
	}

	return req.Name, nil
ERR:
	return "", res.Err()
}

func Search() {

}

//func SelectIanRecord(c *gin.Context, req *params.QueryRequest) (response []*params.QueryResponse, err error) {
//	response = make([]*params.QueryResponse, 0)
//
//	filter := bson.M{
//		"name": req.Name,
//	}
//	if req.Name == "" {
//		filter = bson.M{}
//	}
//	curs, err := storage.GlobalMongo.DB.Collection("ian").Find(context.TODO(), filter)
//	if curs.Err() != nil {
//		logrus.Errorf("查询日常数据失败 %s", curs.Err())
//		goto ERR
//	}
//	if err = curs.All(context.TODO(), &response); err != nil {
//		goto ERR
//	}
//	return response, nil
//ERR:
//	return nil, err
//}

func SelectIanRecord(c *gin.Context, req *V1.QueryRequest) (response []*V1.QueryResponse, err error) {
	var (
		bq      = elastic.NewBoolQuery()
		eq      = elastic.NewExistsQuery("name") // 排除无效脚本
		queries []elastic.Query
		client  *elastic.Client
		daoRes  *elastic.SearchResult
		logger  = logger2.NewLogger()
	)
	client, err = clients.NewESClient()
	defer func() { client.CloseIndex("ian") }()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return nil, err
	}
	if req.Name != "" {
		queries = append(queries, elastic.NewTermQuery("name", req.Name))
	}
	bq.Must(queries...)
	daoRes, err = client.Search().Index("ian").Query(eq).Query(bq).Do(c)
	if err != nil {
		logger.Error(fmt.Sprintf("请求es失败 : %s", err))
		goto ERR
	}
	for _, hit := range daoRes.Hits.Hits {
		var ephemeralIan model.Ian
		err = json.Unmarshal(hit.Source, &ephemeralIan)
		if err != nil {
			goto ERR
		}
		response = append(response, &V1.QueryResponse{
			Name:       ephemeralIan.Name,
			CreateTime: time.Unix(ephemeralIan.Time.T, 0).Format(time.RFC3339),
			UpdateTime: time.Unix(ephemeralIan.Time.I, 0).Format(time.RFC3339),
			Body: struct {
				Weight float32 `json:"weight"`
			}{
				Weight: ephemeralIan.Body.Weight,
			},
			BETre: struct {
				Core       int `json:"core"`
				Runner     int `json:"runner"`
				Support    int `json:"support"`
				Squat      int `json:"squat"`
				EasyBurpee int `json:"easy_burpee"`
				Chair      int `json:"chair"`
				Stretch    int `json:"stretch"`
			}{
				Core:       ephemeralIan.BETre.Core,
				Runner:     ephemeralIan.BETre.Runner,
				Support:    ephemeralIan.BETre.Support,
				Squat:      ephemeralIan.BETre.Squat,
				EasyBurpee: ephemeralIan.BETre.EasyBurpee,
				Chair:      ephemeralIan.BETre.Chair,
				Stretch:    ephemeralIan.BETre.Stretch,
			},
			Worker: struct {
				Vol1 string `json:"vol1"`
				Vol2 string `json:"vol2"`
				Vol3 string `json:"vol3"`
				Vol4 string `json:"vol4"`
			}{
				Vol1: ephemeralIan.Worker.Vol1,
				Vol2: ephemeralIan.Worker.Vol2,
				Vol3: ephemeralIan.Worker.Vol3,
				Vol4: ephemeralIan.Worker.Vol4,
			},
		})

	}

	return response, nil
ERR:
	return nil, err
}

//func AppendIanRecord(c *gin.Context, req *params.AppendRequestInfo) (*model.Ian, error) {
//	var (
//		ian = model.Ian{}
//		err error
//	)
//
//	filter := bson.M{
//		"name": req.Name,
//	}
//	res := storage.GlobalMongo.DB.Collection("ian").FindOne(context.TODO(), filter)
//	if res.Err() != nil {
//		logrus.Errorf("查询日常数据失败 %s", res.Err())
//		goto ERR
//	}
//	if err := res.Decode(&ian); err != nil {
//		goto ERR
//	}
//	ian.BETre.Core += req.BETre.Core
//	ian.BETre.Runner += req.BETre.Runner
//	ian.BETre.Support += req.BETre.Support
//	ian.BETre.Squat += req.BETre.Squat
//	ian.BETre.EasyBurpee += req.BETre.EasyBurpee
//	ian.BETre.Chair += req.BETre.Chair
//	ian.BETre.Stretch += req.BETre.Stretch
//	ian.Worker.Vol1 = fmt.Sprintf("%s;%s", ian.Worker.Vol1, req.Worker.Vol1)
//	ian.Worker.Vol2 = fmt.Sprintf("%s;%s", ian.Worker.Vol2, req.Worker.Vol2)
//	ian.Worker.Vol3 = fmt.Sprintf("%s;%s", ian.Worker.Vol3, req.Worker.Vol3)
//	ian.Worker.Vol4 = fmt.Sprintf("%s;%s", ian.Worker.Vol4, req.Worker.Vol4)
//
//	if storage.GlobalMongo.DB.Collection("ian").FindOneAndReplace(context.TODO(), filter, &ian).Err() != nil {
//		goto ERR
//	}
//	return &ian, nil
//ERR:
//	return nil, err
//}

func AppendIanRecord(c *gin.Context, req *V1.AppendRequestInfo) (*model.Ian, error) {
	var (
		bq      = elastic.NewBoolQuery()
		queries []elastic.Query
		client  *elastic.Client
		daoRes  *elastic.SearchResult
		ian     model.Ian
		err     error
		logger  = logger2.NewLogger()
	)
	client, err = clients.NewESClient()
	if err != nil {
		logger.Error(fmt.Sprintf("初始化 es 失败 %s", err))
		return nil, err
	}
	queries = append(queries, elastic.NewTermQuery("id", req.Id))
	bq.Must(queries...)
	daoRes, err = client.Search("ian").Query(bq).Size(1).Do(c)
	if len(daoRes.Hits.Hits) != 1 {
		goto ERR
	}
	if err := json.Unmarshal(daoRes.Hits.Hits[0].Source, &ian); err != nil {
		goto ERR
	}
	ian.Worker.Vol1 = fmt.Sprintf("%s;%s", ian.Worker.Vol1, req.Worker.Vol1)
	ian.Worker.Vol2 = fmt.Sprintf("%s;%s", ian.Worker.Vol2, req.Worker.Vol2)
	ian.Worker.Vol3 = fmt.Sprintf("%s;%s", ian.Worker.Vol3, req.Worker.Vol3)
	ian.Worker.Vol4 = fmt.Sprintf("%s;%s", ian.Worker.Vol4, req.Worker.Vol4)

	return &ian, err

ERR:
	return nil, nil
}

func CreateIanRecordV2(c *gin.Context, req *V1.CreateIanRecordRequest) (*V1.CreateIanRecordResponse, error) {
	var (
		err error
	)
	tRecord := mysql.TRecord{
		Name:       req.Name,
		Weight:     req.Weight,
		BF:         req.BF,
		LUN:        req.LUN,
		DIN:        req.DIN,
		EXTRA:      req.EXTRA,
		Core:       req.Core,
		Runner:     req.Runner,
		Support:    req.Support,
		Squat:      req.Squat,
		EasyBurpee: req.EasyBurpee,
		Chair:      req.Chair,
		Stretch:    req.Stretch,
		Vol1:       req.Vol1,
		Vol2:       req.Vol2,
		Vol3:       req.Vol3,
		Vol4:       req.Vol4,
		Content:    req.Content,
		Region:     req.Region,
	}
	db := clients.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[req.Region])
	res, _, err := mysql.Create(db.Client, &tRecord)
	if err != nil {
		logrus.Errorf("create record failed %s: %s", err, res)
		return &V1.CreateIanRecordResponse{}, err
	}
	return &V1.CreateIanRecordResponse{
		Id: tRecord.ID,
	}, nil

}

func SelectIanRecordsV2(c *gin.Context, region string, name string, startTime, modifyTime int64, limit int) (*V1.SelectIanRecordResponse, error) {
	var (
		err     error
		records []*mysql.TRecord
		where   = "1 = 1"
		res     = &V1.SelectIanRecordResponse{
			Items: []V1.IanRecordInfo{},
		}
	)
	if name != "" {
		where = fmt.Sprintf("%s and name = '%s'", where, name)
	}
	if startTime != 0 {
		where = fmt.Sprintf("%s and create_time >= %d", where, startTime)
	}
	if modifyTime != 0 {
		where = fmt.Sprintf("%s and update_time >= %d", where, modifyTime)
	}

	db := clients.NewMysqlConn(config.Conf.Backend.MysqlConfig.Clusters[region])
	result, _, err := mysql.GetValueByRaw(db.Client, &records, "t_record", where)
	if err != nil {
		logrus.Errorf("create record failed %s: %s", err, result)
		return nil, err
	}
	for _, record := range records {
		res.Items = append(res.Items, V1.IanRecordInfo{
			Id:         record.ID,
			CreateTime: record.CreateTime,
			ModifyTime: record.ModifyTime,
			Name:       record.Name,
			Weight:     record.Weight,
			BF:         record.BF,
			LUN:        record.LUN,
			DIN:        record.DIN,
			EXTRA:      record.EXTRA,
			Core:       record.Core,
			Runner:     record.Runner,
			Support:    record.Support,
			Squat:      record.Squat,
			EasyBurpee: record.EasyBurpee,
			Chair:      record.Chair,
			Stretch:    record.Stretch,
			Vol1:       record.Vol1,
			Vol2:       record.Vol2,
			Vol3:       record.Vol3,
			Vol4:       record.Vol4,
			Content:    record.Content,
			Extra:      record.Extra,
			Region:     record.Region,
		})
	}

	return res, nil

}

//func AddDayForm(c *gin.Context) {
//	var (
//		ian model.ShadowPriest
//		err error
//	)
//
//	err = c.ShouldBindJSON(&ian)
//	if err != nil {
//		logrus.Errorf("%s", err)
//		baseHandlers.RenderData(c, nil, err)
//		return
//	}
//	err = storage.GloablMongo.C("ian").Insert(ian)
//	if err != nil {
//		logrus.Errorf("%s", err)
//		baseHandlers.RenderData(c, nil, err)
//		return
//	}
//	baseHandlers.RenderData(c, nil, nil)
//
//}
//
//func UpdateForm(c *gin.Context) {
//	var (
//		ian model.ShadowPriest
//		err error
//	)
//	if err := utils.EnsureJson(c, &ian); err != nil {
//		baseHandlers.RenderData(c, nil, err)
//		return
//	}
//
//	err = storage.GloablMongo.C("ian").Update(bson.M{
//		"id": ian.Id,
//	}, bson.M{
//		"$set": utils.ConvBsonNoErr(ian),
//	})
//	if err != nil {
//		logrus.Errorf("%s", err)
//	}
//	baseHandlers.RenderData(c, nil, err)
//
//}
//
//func SelectForm(c *gin.Context) {
//	var (
//		ian       model.ShadowPriest
//		queryInfo model.ShadowPriestQueryRequestInfo
//		err       error
//		selector  bson.M
//	)
//	if err = utils.EnsureJson(c, &queryInfo); err != nil {
//		goto ERR
//	}
//	if err = bson.UnmarshalJSON([]byte(queryInfo.Query), &selector); err != nil {
//		goto ERR
//	}
//	if err = storage.GloablMongo.C("ian").Find(&selector).One(&ian); err != nil {
//		goto ERR
//	}
//
//	baseHandlers.RenderData(c, ian, err)
//	return
//ERR:
//	baseHandlers.RenderData(c, nil, err)
//	return
//
//}
