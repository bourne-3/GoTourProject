package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func main() {

	// way1
	//r := gin.Default()
	//// 路由注册
	//r.GET("/index", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "hello, world"})
	//})
	//r.Run()

	// way2
	//router := routers.NewRouter()
	//s := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()

	// way3
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}

// 在 main 方法之前自动执行，它的执行顺序是：全局变量初始化 =》init 方法 =》main 方法，
func init() {
	// 初始化流程控制
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil

}
