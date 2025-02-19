package server

import (
	"blog1/src/model/gorm_model"
	"blog1/src/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BackAllData(c *gin.Context) {
	categoryName := c.Param("category_name")

	backAllData, err := svc.BackAllDateService(categoryName)
	if err != nil {
		zap.L().Error("svc.BackAllDateService() err: ", zap.Error(err))
		response.Fail(c, "back_index", response.ErrServer)
		return
	}

	response.Success(c, "back_index", backAllData)
	return
}

func BackNewPost(c *gin.Context) {
	// 接收请求
	postModel := gorm_model.Post{}
	err := c.ShouldBind(&postModel)
	if err != nil {
		zap.L().Error("c.ShouldBind(&postModel) err: ", zap.Error(err))
		response.Fail(c, "new_post", response.ErrServer)
		return
	}

	// 获取上传的文件
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		zap.L().Error("c.Request().FormFile() err: ", zap.Error(err))
		response.Fail(c, "new_post", response.ErrServer)
		return
	}
	defer file.Close()

	err = svc.NewPostService(postModel, file, fileHeader)
	if err != nil {
		zap.L().Error("svc.BackAllDateService() err: ", zap.Error(err))
		response.Fail(c, "new_post", response.ErrServer)
		return
	}

	backAllData, err := svc.BackAllDateService("")
	if err != nil {
		zap.L().Error("NewPost() svc.BackAllDateService() err: ", zap.Error(err))
		response.Fail(c, "new_post", response.ErrServer)
		return
	}

	response.SuccessWithMsg(c, "new_post", "发布成功", backAllData)
}

func NewPostView(c *gin.Context) {
	backAllData, err := svc.BackAllDateService("")
	if err != nil {
		zap.L().Error("NewPostView() svc.BackAllDateService() err: ", zap.Error(err))
		response.Fail(c, "new_post", response.ErrServer)
		return
	}

	response.Success(c, "new_post", backAllData)
}

func BackPost(c *gin.Context) {
	slug := c.Param("slug")

	backPostModel, err := svc.GetBackPostService(slug)
	if err != nil {

		fmt.Println("svc.GetBackPostService() err: ", zap.Error(err))
		return
	}

	response.Success(c, "back_post", backPostModel)
}

func BackDelPost(c *gin.Context) {
	slug := c.Param("slug")

	err := svc.BackDelPostService(slug)
	if err != nil {
		zap.L().Error("svc.BackDelPostService() err: ", zap.Error(err))
		response.Fail(c, "back_post", response.ErrServer)
		return
	}

	response.SuccessWithMsg(c, "back_post", "删除成功", "")

}
