package main

import (
	"fmt"
	"layout/pkg/setting"
	"layout/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	serve := &http.Server{
		// 监听的TCP地址，格式为 :8000
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		// http句柄，实质为ServerHTTP，用于处理程序的响应HTTP请求
		Handler: router,
		// 允许读取的最大时间
		ReadTimeout: setting.ReadTimeout,
		// 允许写入的最大时间
		WriteTimeout: setting.WriteTimeout,
		// 请求头的最大字节数
		MaxHeaderBytes: 1 << 20,
	}

	err := serve.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
