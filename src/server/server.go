package server

import (
	"github.com/gin-gonic/gin"
	"staging/src/pkg/settings"
	"staging/src/service"
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
