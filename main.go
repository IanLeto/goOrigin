package main

import (
	"github.com/gin-gonic/gin"
	"goOrigin/router"
	"net/http"
)

func main() {
	g := gin.New()
	router.Load(g, nil)
	go func() {

	}()
	http.ListenAndServe(":8080", g)
}
