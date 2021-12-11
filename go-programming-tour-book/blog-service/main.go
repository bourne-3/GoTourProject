package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 路由注册
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello, world"})
	})
	r.Run()
}
