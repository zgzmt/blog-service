package main

import (
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	setting2 "blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)
func init(){
	err := setupSetting()
	if err != nil {
		log.Fatal("init.setupSetting err:%v",err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatal("init.setupSetting err:%v",err)
	}
}

func main(){
	//r := gin.Default()
	//
	//r.GET("/hello", func(context *gin.Context) {
	//	context.JSON(200,gin.H{"message":"hello"})
	//})
	//r.Run()
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error{
	setting,err := setting2.NewString()
	if err != nil{
		return err
	}
	err = setting.ReadSection("Server",&global.ServerSetting)
	if err != nil{
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine()error{
	var err error
	global.DBEngine ,err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil{
		return err
	}
	return nil
}