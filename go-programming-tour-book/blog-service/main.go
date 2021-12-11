package main

import (
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"net/http"
	"time"
)

func main() {
	//r := gin.Default()
	//// 路由注册
	//r.GET("/index", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "hello, world"})
	//})
	//r.Run()

	// way2
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
