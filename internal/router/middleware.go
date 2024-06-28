package router

import (
	"context"
)

func AuthMiddle(ctx context.Context, region string) (bool, error) {
	//var (
	//	err    error
	//	logger = logger2.NewLogger()
	//	client *redis.ClusterClient
	//	allow  bool
	//	token  string
	//	//result map[string]string
	//	//user   entity.User
	//)
	//token, err = conv.String(token)
	//utils.NoError(err)
	//client = redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs: []string{},
	//})
	//result, err = client.HGetAll(token).Result()
	//switch err {
	//case redis.Nil:
	//	userStr := entity.UserStr(token)
	//	user = &userStr
	//
	//	//user = &entity.UserStr(token)
	//}
	//user = &entity.UserRedis{}
	return false, nil

}
