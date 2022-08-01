package params

import "github.com/gin-gonic/gin"

type BaseResponseInfo struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

type BaseErrResponseInfo struct {
	*BaseResponseInfo
}

type BaseK8sRequestInfo struct {
	Selector string
}

func BuildErrResponse(c *gin.Context, info *BaseErrResponseInfo) {
	c.JSON(400, gin.H{
		"Data":    nil,
		"Message": info.Message,
		"Code":    info.Code,
	})
}

func BuildResponse(c *gin.Context, info *BaseResponseInfo) {
	c.JSON(200, gin.H{
		"Data":    info.Data,
		"Message": info.Message,
		"Code":    200,
	})
}

func BuildErrInfo(code int, msg string) *BaseErrResponseInfo {
	return &BaseErrResponseInfo{
		&BaseResponseInfo{
			Data:    nil,
			Message: msg,
			Code:    code,
		},
	}
}

func BuildInfo(data interface{}) *BaseResponseInfo {
	return &BaseResponseInfo{
		Data:    data,
		Message: "Success",
		Code:    200,
	}
}
