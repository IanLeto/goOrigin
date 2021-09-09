package ianHandlers

import (
	"github.com/gin-gonic/gin"
	"goOrigin/model"
)

func AddDayForm(c *gin.Context) {
	var (
		ian model.IanUI
	)
	err := c.ShouldBindJSON(ian)
	if err != nil {

	}

}
