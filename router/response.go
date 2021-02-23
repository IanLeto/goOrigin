package router

import "github.com/gin-gonic/gin"

func renderData(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(200, gin.H{"msg": err, "data": nil})
	}

}
