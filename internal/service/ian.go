package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
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
	return res.InsertedID, nil
ERR:
	return "", nil
}

func DeleteIanRecord(c *gin.Context, id string) (delCount int64, err error) {
	res, err := storage.Mongo.DB.Collection("ian").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		logrus.Errorf("删除日常数据失败 %s", err)
		goto ERR
	}
	return res.DeletedCount, nil
ERR:
	return 0, nil

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
