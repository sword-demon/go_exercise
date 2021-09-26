package routers

import (
	"github.com/gin-gonic/gin"
	"layout/pkg/setting"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 先设置运行模式
	gin.SetMode(setting.RunMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 路由组


	return router
}