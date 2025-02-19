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

	// 后台文章
	bpRouter := backRouter.Group("/posts")
	bpRouter.POST("/new", BackNewPost)            // 增                                                                                                                                                                                                                                                            `
	bpRouter.DELETE("/:slug", BackDelPost)        // 删
	bpRouter.PUT("/:slug", BackEditPost)          // 改
	bpRouter.GET("/", BackAllData)                // 查所有
	bpRouter.GET("/:slug", BackPost)              // 查单篇
	bpRouter.GET("/new", BackNewPostView)         // 增文章页面
	bpRouter.GET("/edit/:slug", BackEditPostView) // 编辑文章页面

	// 后台分类
	bcRouter := backRouter.Group("/categories")
	bcRouter.GET("/", BackCategoriesView)               // 查所有
	bcRouter.DELETE("/:categoryId", BackCategoriesView) // 查所有
	bcRouter.PUT("/:categoryId", BackCategoriesView)    // 查所有

	// 后台评论
	bcoRouter := backRouter.Group("/comments")
	bcoRouter.DELETE("/:comment_id", BackDelComment) // 删

	return r
}
