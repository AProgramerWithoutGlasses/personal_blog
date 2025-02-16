package server

import (
	"blog1/src/model/gorm_model"
	"blog1/src/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAllBackData(c *gin.Context) {
	categoryName := c.Param("category_name")

	backAllData, err := svc.BackAllDateService(categoryName)
	if err != nil {
		zap.L().Error("svc.BackAllDateService() err: ", zap.Error(err))
		response.Fail(c, "index_back", response.ErrServer)
		return
	}

	response.Success(c, "index_back", backAllData)
	return
}

func NewPost(c *gin.Context) {
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

func GetPostBack(c *gin.Context) {
	//slug := c.Param("slug")
	//
	//GetBackPostService
	//
	//response.Success(c, "new_post", backAllData)
}
