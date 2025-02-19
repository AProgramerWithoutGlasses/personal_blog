package server

import (
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

	//backCategoriesModel, err := svc.BackEditCategoryService()
	//if err != nil {
	//	zap.L().Error("BackCategoriesView() svc.BackCategoriesService() err: ", zap.Error(err))
	//	response.Fail(c, "back_categories", response.ErrServer)
	//	return
	//}
	//
	//response.Success(c, "back_categories", backCategoriesModel)
}
