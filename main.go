package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goOrigin/config"
	"goOrigin/define"
	"goOrigin/event"
	_ "goOrigin/init"
	"goOrigin/router"
	"goOrigin/utils"
	"net/http"
)

func main() {
	// 让包主动注册进来 并不符合golang 的设计规范
	// but 符合人的思维

	for _, handler := range define.InitHandler {
		if err := handler(); err != nil {
			panic(err)
		}
	}
	g := gin.New()
	gin.SetMode(config.GlobalConfig.RunMode)
	router.Load(g, nil)
	event.GlobalEventBus.PubPeriodicTask("ccPing", nil)
	http.ListenAndServe(":"+utils.ConvOrDefaultString(viper.Get("port"), "8080"), g)

}
