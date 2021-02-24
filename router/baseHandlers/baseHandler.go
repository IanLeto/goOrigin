package baseHandlers

import (
	"github.com/gin-gonic/gin"
	"goOrigin/errors"
	"net/http"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// 旧版本
func RenderData(c *gin.Context, data interface{}, err error) {
	c.JSON(http.StatusOK, http.Response{

	})
}

func RenderResponse(c *gin.Context, data interface{}, err *errors.Err) {
	if err == nil {
		err = &errors.Err{
			Code:    0,
			Message: "OK",
			Err:     nil,
		}
	}
	c.JSON(err.Code, Response{
		Msg:  err.Message,
		Code: err.Code,
		Data: data,
	})
}
