package server

import (
	"blog1/src/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func indexView(name string) gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := svc.IndexService()
		if err != nil {
			fmt.Println(err)
			return
		}

		response.Success(c, name, res)
	}

}

func postView(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 动态路由
		slug := c.Param("slug")

		postModel, err := svc.PostService(slug)
		if err != nil {
			fmt.Println("svc.PostService() err: ", err)
			return
		}

		response.Success(c, name, postModel)
	}

}

func categoryView(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 动态路由
		categoryName := c.Param("name")

		categoryModel, err := svc.CategoryService(categoryName)
		if err != nil {
			fmt.Println("svc.PostService() err: ", err)
			return
		}

		response.Success(c, name, categoryModel)
	}
}
