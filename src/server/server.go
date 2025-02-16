package server

import (
	"blog1/src/pkg/settings"
	"blog1/src/service"
	"github.com/gin-gonic/gin"
)

var svc *service.Service

//func initRouter() *gin.Engine {
//	// 创建一个默认的路由引擎
//	r := gin.Default()
//	r.Use(logger.GinLogger(), logger.GinRecovery(true))
//	return r
//}

func Init(app *settings.AppConfig) *gin.Engine {
	svc = service.InitService(app)
	return InitRouter()
}

// 前台

// 后台
