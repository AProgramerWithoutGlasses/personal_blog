package server

import (
	"blog1/src/logger"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 预加载html模板
	r.LoadHTMLGlob("static/views/*")

	// 加载静态资源
	r.Static("/static", "./static")

	r.POST("/login", loginFrontHandler)
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
	bpRouter := backRouter.Group("/posts")
	bpRouter.POST("/new", BackNewPost)     // 增                                                                                                                                                                                                                                                            `
	bpRouter.DELETE("/:slug", BackDelPost) // 删
	bpRouter.PUT("/:slug")                 // 改
	bpRouter.GET("/", BackAllData)         // 查所有
	bpRouter.GET("/:slug", BackPost)       // 查单篇
	bpRouter.GET("/new", NewPostView)      //

	bcRouter := backRouter.Group("/categories")
	bcRouter.GET("/:category_name", BackAllData) // 查单篇

	return r
}
