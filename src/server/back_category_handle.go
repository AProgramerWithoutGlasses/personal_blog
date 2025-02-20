package server

import (
	"blog1/src/model"
	"blog1/src/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BackCategoriesView(c *gin.Context) {
	backCategoriesModel, err := svc.BackCategoriesService()
	if err != nil {
		zap.L().Error("BackCategoriesView() svc.BackCategoriesService() err: ", zap.Error(err))
		response.Fail(c, "back_categories", response.ErrServer)
		return
	}

	response.Success(c, "back_categories", backCategoriesModel)
}

func BackEditCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")

	backEditCategoryModel := model.BackEditCategoriesModel{}
	err := c.ShouldBind(&backEditCategoryModel)
	if err != nil {
		zap.L().Error("BackEditCategory() ShouldBind() err: ", zap.Error(err))
		response.Fail(c, "back_edit_categories", response.ErrServer)
	}

	err = svc.BackEditCategoryService(categoryId, backEditCategoryModel.Name)
	if err != nil {
		zap.L().Error("BackEditCategory() svc.BackEditCategoryService() err: ", zap.Error(err))
		response.Fail(c, "back_categories", response.ErrServer)
		return
	}

	response.SuccessWithMsg(c, "back_categories", "修改成功", "")
}

func BackDelCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")

	err := svc.BackDelCategoryService(categoryId)
	if err != nil {
		zap.L().Error("BackDelCategory() svc.BackDelCategoryService() err: ", zap.Error(err))
		response.Fail(c, "back_categories", response.ErrServer)
		return
	}

	response.SuccessWithMsg(c, "back_categories", "删除成功", "")
}

func BackNewCategory(c *gin.Context) {
	backNewCategoriesModel := model.BackNewCategoriesModel{}
	err := c.ShouldBind(&backNewCategoriesModel)
	if err != nil {
		zap.L().Error("BackNewCategory() ShouldBind() err: ", zap.Error(err))
		response.Fail(c, "back_categories", response.ErrServer)
		return
	}

	err = svc.BackNewCategoryService(backNewCategoriesModel)
	if err != nil {
		zap.L().Error("BackDelCategory() svc.BackNewCategoryService() err: ", zap.Error(err))
		response.Fail(c, "back_categories", response.ErrServer)
		return
	}

	response.SuccessWithMsg(c, "back_categories", "新增成功", "")
}
