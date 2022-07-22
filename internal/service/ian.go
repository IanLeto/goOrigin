package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goOrigin/internal/model"
	"goOrigin/internal/params"
	"goOrigin/pkg/storage"
)

func CreateIanRecord(c *gin.Context, req params.CreateIanRequestInfo) (id interface{}, err error) {
	var (
		ian = model.NewIan(req)
	)
	res, err := storage.Mongo.DB.Collection("ian").InsertOne(context.TODO(), &ian)
	if err != nil {
		logrus.Errorf("创建日常数据失败")
		goto ERR
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
ERR:
	return "", nil
}

func DeleteIanRecord(c *gin.Context, id string) (delCount int64, err error) {

	res, err := storage.Mongo.DB.Collection("ian").DeleteMany(context.TODO(), bson.M{"name": id})
	if err != nil {
		logrus.Errorf("删除日常数据失败 %s", err)
		goto ERR
	}
	return res.DeletedCount, nil
ERR:
	return 0, nil

}

func UpdateIanRecord(c *gin.Context, req params.CreateIanRequestInfo) (id interface{}, err error) {
	var (
		ian = model.NewIan(req)
	)

	res := storage.Mongo.DB.Collection("ian").FindOneAndReplace(context.TODO(), bson.M{"name": req.Name},
		&ian)

	if res.Err() != nil {
		logrus.Errorf("创建日常数据失败 %s", res.Err())
		goto ERR
	}
	return req.Name, nil
ERR:
	return "", res.Err()
}

// SelectIanRecordDetail 详细过滤的demo
func SelectIanRecordDetail() {
	
}

func SelectIanRecord(c *gin.Context, req *params.QueryRequest) (response []*params.QueryResponse, err error) {
	response = make([]*params.QueryResponse, 0)

	filter := bson.M{
		"name": req.Name,
	}
	if req.Name == "" {
		filter = bson.M{}
	}
	curs, err := storage.Mongo.DB.Collection("ian").Find(context.TODO(), filter)
	if curs.Err() != nil {
		logrus.Errorf("查询日常数据失败 %s", curs.Err())
		goto ERR
	}
	if err = curs.All(context.TODO(), &response); err != nil {
		goto ERR
	}
	return response, nil
ERR:
	return nil, err
}

func AppendIanRecord(c *gin.Context, req *params.AppendRequestInfo) (*model.Ian, error) {
	var (
		ian = model.Ian{}
		err error
	)

	filter := bson.M{
		"name": req.Name,
	}
	res := storage.Mongo.DB.Collection("ian").FindOne(context.TODO(), filter)
	if res.Err() != nil {
		logrus.Errorf("查询日常数据失败 %s", res.Err())
		goto ERR
	}
	if err := res.Decode(&ian); err != nil {
		goto ERR
	}
	ian.BETre.Core += req.BETre.Core
	ian.BETre.Runner += req.BETre.Runner
	ian.BETre.Support += req.BETre.Support
	ian.BETre.Squat += req.BETre.Squat
	ian.BETre.EasyBurpee += req.BETre.EasyBurpee
	ian.BETre.Chair += req.BETre.Chair
	ian.BETre.Stretch += req.BETre.Stretch
	ian.Worker.Vol1 = fmt.Sprintf("%s;%s", ian.Worker.Vol1, req.Worker.Vol1)
	ian.Worker.Vol2 = fmt.Sprintf("%s;%s", ian.Worker.Vol2, req.Worker.Vol2)
	ian.Worker.Vol3 = fmt.Sprintf("%s;%s", ian.Worker.Vol3, req.Worker.Vol3)
	ian.Worker.Vol4 = fmt.Sprintf("%s;%s", ian.Worker.Vol4, req.Worker.Vol4)

	if storage.Mongo.DB.Collection("ian").FindOneAndReplace(context.TODO(), filter, &ian).Err() != nil {
		goto ERR
	}
	return &ian, nil
ERR:
	return nil, err
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
//	err = storage.Mongo.C("ian").Insert(ian)
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
//	err = storage.Mongo.C("ian").Update(bson.M{
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
//	if err = storage.Mongo.C("ian").Find(&selector).One(&ian); err != nil {
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