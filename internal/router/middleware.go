package router

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"goOrigin/internal/model/entity"
	logger2 "goOrigin/pkg/logger"
	"goOrigin/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {

	//user = &entity.UserRedis{}
	return func(c *gin.Context) {
		var (
			err      error
			logger   = logger2.Logger
			client   *redis.ClusterClient
			allow    bool
			token    string
			loginUrl string
			user     entity.User
		)

		token = c.GetHeader("token")
		if token == "" {
			goto noAuth
		}
		loginUrl, err = conv.String(c.Value("loginUrl"))
		utils.NoError(err)
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{},
		})
		_, err = client.HGetAll(token).Result()
		logger.Info(err.Error())
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
		if !allow {
			goto noAuth
		} else {
			c.Set("user", user)
			c.Next()
		}

	noAuth:

		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return

	}

}
