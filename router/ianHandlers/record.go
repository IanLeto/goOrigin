package ianHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"goOrigin/model"
	"goOrigin/router/baseHandlers"
	"goOrigin/storage"
)

func AddDayForm(c *gin.Context) {
	var (
		ian model.IanUI
		err error
	)

	err = c.ShouldBindJSON(&ian)
	if err != nil {
		baseHandlers.RenderData(c, nil, err)
		return
	}
	err = storage.Mongo.C("ian").Insert(ian)
	if err != nil {
		baseHandlers.RenderData(c, nil, err)
	}
	baseHandlers.RenderData(c, nil, nil)

}

func UpdateForm(c *gin.Context) {
	var (
		ian model.IanUI
		err error
	)
	err = c.ShouldBindJSON(&ian)
	if err != nil {
		logrus.Errorf("%s", err)
	}
	if ian.Weight != 0 {
		info, err := storage.Mongo.C("ian").UpdateAll(bson.M{
			"id": "1",
		}, bson.M{
			"$set": bson.M{
				"weight": 120,
			},
		})
		logrus.Errorf("%s", err)
		baseHandlers.RenderResponse(c, info, nil)

	}

}
