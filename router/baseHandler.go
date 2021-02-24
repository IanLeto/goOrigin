package router

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

func renderData(c *gin.Context, data interface{}, err error) {
	c.JSON(http.StatusOK, http.Response{

	})
}

func RenderResponse(c *gin.Context, data interface{}, err *errors.Errno) {
	if err == nil {
		err = &errors.Errno{
			Code:    0,
			Message: "OK",
		}
	}
	c.JSON(err.Code, Response{
		Msg:  err.Message,
		Code: err.Code,
		Data: data,
	})
}
