package baseHandlers

import (
	"github.com/gin-gonic/gin"
	"goOrigin/errors"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func RenderData(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(200, gin.H{
			"data": data,
			"err":  err,
		})
	}
}

func renderMessage(c *gin.Context, v interface{}) {
	if v == nil {
		c.JSON(200, gin.H{
			"err": "",
		})
	}
	switch t := v.(type) {
	case string:
		c.JSON(200, gin.H{
			"err": t,
		})
	case error:
		c.JSON(200, gin.H{
			"err": t.Error(),
		})
	}
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
