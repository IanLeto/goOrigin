package router

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err      error
			allow    bool
			token    string
			loginUrl string
			user     entity.User
		)

		token = c.GetHeader("Authorization")
		if token == "" {
			goto noAuth
		}
		loginUrl, err = conv.String(c.Value("loginUrl"))
		utils.NoError(err)
		switch err {
		case redis.Nil:
			userStr := entity.UserStr(token)
			user = &userStr
			allow, err = user.Auth(token, loginUrl)
			utils.NoError(err)
		default:
			fmt.Println("实验性代码")
			userStr := entity.UserStr(token)
			user = &userStr
			allow, err = user.Auth(token, loginUrl)
			utils.NoError(err)
			//user = &entity.UserStr(token)
		}
		fmt.Println(c.Request.URL.String())
		if !allow {
			goto noAuth
		} else {
			c.Set("user", user)
			c.Next()
			return
		}

	noAuth:

		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return

	}

}
