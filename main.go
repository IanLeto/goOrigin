package main

import (
	"github.com/gin-gonic/gin"
	"goOrigin/define"
	"goOrigin/event"
	_ "goOrigin/init"
	"goOrigin/router"
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
	gin.SetMode(gin.ReleaseMode)
	router.Load(g, nil)
	go func() {

	}()
	event.GlobalEventBus.PubPeriodicTask("ccPing", nil)
	//http.ListenAndServe(":"+utils.ConvOrDefaultString(viper.Get("port"), "8080"), g)
	http.ListenAndServe("localhost:8080", g)

}
