package ianHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"goOrigin/model"
	"goOrigin/router/baseHandlers"
	"goOrigin/storage"
	"goOrigin/utils"
)

func AddDayForm(c *gin.Context) {
	var (
		ian model.ShadowPriest
		err error
	)

	err = c.ShouldBindJSON(&ian)
	if err != nil {
		logrus.Errorf("%s", err)
		baseHandlers.RenderData(c, nil, err)
		return
	}
	err = storage.Mongo.C("ian").Insert(ian)
	if err != nil {
		logrus.Errorf("%s", err)
		baseHandlers.RenderData(c, nil, err)
		return
	}
	baseHandlers.RenderData(c, nil, nil)

}

func UpdateForm(c *gin.Context) {
	var (
		ian model.ShadowPriest
		err error
	)
	if err := utils.EnsureJson(c, &ian); err != nil {
		baseHandlers.RenderData(c, nil, err)
		return
	}

	err = storage.Mongo.C("ian").Update(bson.M{
		"id": ian.Id,
	}, bson.M{
		"$set": utils.ConvBson(ian),
	})
	if err != nil {
		logrus.Errorf("%s", err)
	}
	baseHandlers.RenderData(c, nil, err)

}
