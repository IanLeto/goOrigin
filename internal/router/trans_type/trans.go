package trans_type

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goOrigin/API/V1"
	"goOrigin/internal/logic"
	"goOrigin/internal/model/entity"
)

func CreateTransInfo(c *gin.Context) {
	var (
		req = &V1.CreateTransInfoReq{}
		res = &V1.CreateTransInfoResponse{}
		err error
	)
	if req.Region == "" {
		c.Set("region", "win")
	} else {
		c.Set("region", req.Region)
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	err = logic.CreateType(c, req.Region, req.Items)
	if err != nil {
		logrus.Errorf("%s", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(res))
	return
ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("create recoed failed by %s", err)))
}

func GetTransInfoList(c *gin.Context) {
	var (
		req   = &V1.GetTransInfoListReq{}      // 请求结构体
		res   = &V1.GetTransInfoListResponse{} // 响应结构体
		err   error
		total int64
	)
	// 调用逻辑层查询函数
	var list []*entity.TransInfoEntity

	// 设置默认 region
	if req.Region == "" {
		c.Set("region", "win")
	} else {
		c.Set("region", req.Region)
	}

	// 绑定 JSON 请求体
	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("failed to bind JSON: %v", err)
		goto ERR
	}

	list, total, err = logic.QueryTransTypeList(c, req.Region, req.Project, req.TransType, req.Page, req.PageSize)
	if err != nil {
		logrus.Errorf("query logic failed: %v", err)
		goto ERR
	}

	res.Items = list
	res.Total = total
	res.Page = req.Page
	res.PageSize = req.PageSize
	V1.BuildResponse(c, V1.BuildInfo(res))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("query trans info failed: %s", err)))
}

func DeleteTransInfo(c *gin.Context) {
	var (
		req = &V1.DeleteTransInfoReq{}
		err error
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("bind json failed: %v", err)
		goto ERR
	}

	if req.Region == "" {
		req.Region = "win"
	}

	err = logic.DeleteTransInfo(c, req.Region, req.Project, req.TransType)
	if err != nil {
		logrus.Errorf("delete trans info failed: %v", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(nil))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("delete failed: %s", err)))
}

func UpdateTransInfo(c *gin.Context) {
	var (
		req = &V1.UpdateTransInfoReq{}
		err error
	)
	infoEntity := convertToEntity(req.Item)

	if err = c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("bind json failed: %v", err)
		goto ERR
	}

	if req.Region == "" {
		req.Region = "win"
	}

	// ✨ 转换请求结构体为 infoEntity

	err = logic.UpdateTransInfo(c, req.Region, infoEntity)
	if err != nil {
		logrus.Errorf("update failed: %v", err)
		goto ERR
	}

	V1.BuildResponse(c, V1.BuildInfo(nil))
	return

ERR:
	V1.BuildErrResponse(c, V1.BuildErrInfo(0, fmt.Sprintf("update failed: %s", err)))
}

// 🔁 转换函数：将 API 层结构体转换为 entity 层结构体
func convertToEntity(item *V1.UpdateTransInfo) *entity.TransInfoEntity {
	if item == nil {
		return nil
	}

	var codes []*entity.ReturnCodeEntity
	for _, rc := range item.ReturnCodes {
		codes = append(codes, &entity.ReturnCodeEntity{
			ReturnCode:   rc.ReturnCode,
			ReturnCodeCn: rc.ReturnCodeCn,
			TransType:    rc.TransType,
			ProjectID:    rc.Project,
			Status:       rc.Status,
		})
	}

	return &entity.TransInfoEntity{
		Project:    item.Project,
		TransType:  item.TransType,
		Interval:   item.Interval,
		ReturnCode: codes,
	}
}
