package main

import (
	"github.com/gin-gonic/gin"
	"goOrigin/define"
	"goOrigin/router"
	"net/http"
)

func main() {
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

	//http.ListenAndServe(":"+utils.ConvOrDefaultString(viper.Get("port"), "8080"), g)
	http.ListenAndServe("localhost:8080", g)

}
