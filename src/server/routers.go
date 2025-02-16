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

	// 加载静态资源
	r.Static("/static", "./static")

	r.POST("/login", loginFrontHandler("login.html"))
	r.GET("/login", loginFrontView("login.html"))
	r.POST("/register", registerHandler("register.html"))
	r.GET("/register", registerView("register.html"))

	// 前台
	indexRouter := r.Group("/index")
	indexRouter.GET("", indexView("index.html"))
	indexRouter.GET("/post/:slug", postView("post.html"))
	indexRouter.GET("/category/:name", categoryView("category.html"))

	// 后台
	backRouter := r.Group("/back")
	bpRouter := backRouter.Group("/post")
	bpRouter.POST("/") // 增
	bpRouter.DELETE("/:id")
	bpRouter.PUT("/:id")              // 改
	bpRouter.GET("/", GetAllBackData) //posts

	return r
}
