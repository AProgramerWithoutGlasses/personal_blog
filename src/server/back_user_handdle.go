package server

import (
	"blog1/src/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BackUsersView(c *gin.Context) {
	backUsersModel, err := svc.BackUsersViewService()
	if err != nil {
		zap.L().Error("BackUsersView svc.BackUsersViewService()", zap.Error(err))
		response.Fail(c, "back_users", response.ErrServer)
	}

	response.Success(c, "back_users", backUsersModel)
}

func BackDelUser(c *gin.Context) {
	backUsersModel, err := svc.BackUsersViewService()
	if err != nil {
		zap.L().Error("BackUsersView svc.BackUsersViewService()", zap.Error(err))
		response.Fail(c, "back_users", response.ErrServer)
	}

	response.Success(c, "back_users", backUsersModel)
}

func BackEditUser(c *gin.Context) {
	backUsersModel, err := svc.BackUsersViewService()
	if err != nil {
		zap.L().Error("BackUsersView svc.BackUsersViewService()", zap.Error(err))
		response.Fail(c, "back_users", response.ErrServer)
	}

	response.Success(c, "back_users", backUsersModel)
}
