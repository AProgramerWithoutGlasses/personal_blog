package server

import (
	"github.com/gin-gonic/gin"
	"staging/src/logger"
)

func InitRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 预加载html模板
	r.LoadHTMLGlob("static/views/*")

	r.GET("/test", test)
	r.POST("/login", loginFrontHandler("login.html"))
	r.GET("/login", loginFrontView("login.html"))
	r.POST("/register", registerHandler("register.html"))
	//r.GET("/register", registerView("register.html"))

	return r
}
