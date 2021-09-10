package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func EnsureJson(c *gin.Context, v interface{}) error {
	if err := c.ShouldBindJSON(v); err != nil {
		logrus.Errorf("数据格式不对 %s", err)
		return err
	}
	return nil
}

func get()  {
	
}