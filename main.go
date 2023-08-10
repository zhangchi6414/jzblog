package main

import (
	"github.com/gin-gonic/gin"
	"jzblog/global"
	"jzblog/pkg/setting"
	"jzblog/server/routers"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting.err: %s", err)
	}
}

func setupSetting() error {
	se, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = se.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = se.ReadSection("APP", &global.AppSetting)
	if err != nil {
		return err
	}
	err = se.ReadSection("Datebase", &global.DBSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func main() {
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
