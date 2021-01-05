package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goOrigin/define"
	"goOrigin/router"
	"goOrigin/utils"
	"net/http"
)

func main() {
	for _, handler := range define.InitHandler {
		if err := handler(); err != nil {
			panic(err)
		}
	}
	g := gin.New()
	router.Load(g, nil)
	go func() {

	}()

	http.ListenAndServe(":"+utils.ConvOrDefaultString(viper.Get("port"), ""), g)

}
