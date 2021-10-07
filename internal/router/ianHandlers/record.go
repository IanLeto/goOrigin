package ianHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"goOrigin/internal/model"
	"goOrigin/internal/router/baseHandlers"
	"goOrigin/pkg/storage"
	"goOrigin/pkg/utils"
)


// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /hello [get]
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
		"$set": utils.ConvBsonNoErr(ian),
	})
	if err != nil {
		logrus.Errorf("%s", err)
	}
	baseHandlers.RenderData(c, nil, err)

}

func SelectForm(c *gin.Context) {
	var (
		ian       model.ShadowPriest
		queryInfo model.ShadowPriestQueryRequestInfo
		err       error
		selector  bson.M
	)
	if err = utils.EnsureJson(c, &queryInfo); err != nil {
		goto ERR
	}
	if err = bson.UnmarshalJSON([]byte(queryInfo.Query), &selector); err != nil {
		goto ERR
	}
	if err = storage.Mongo.C("ian").Find(&selector).One(&ian); err != nil {
		goto ERR
	}

	baseHandlers.RenderData(c, ian, err)
	return
ERR:
	baseHandlers.RenderData(c, nil, err)
	return

}
