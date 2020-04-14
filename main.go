package main

import (
	"fmt"
	"go-gin-blog/conf"
	"go-gin-blog/router"
	"net/http"
)

// 项目初始化：系统配置，数据库连接 -> 路由 -> 路由中间件 -> 业务层 -> model 层
func main() {
	r := router.InitRouter()

	serverConf := conf.Setting.Server

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", serverConf.HttpPort),
		Handler:      r,
		ReadTimeout:  serverConf.ReadTimeout,
		WriteTimeout: serverConf.WriteTimeout,
	}
	s.ListenAndServe()
}
