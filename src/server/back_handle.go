package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"staging/src/pkg/response"
)

func GetAllBackData(c *gin.Context) {
	backAllData, err := svc.BackAllDateService()
	if err != nil {
		zap.L().Error("svc.BackAllDateService() err: ", zap.Error(err))
		response.Fail(c, "index_back", response.ErrServer)
		return
	}

	response.Success(c, "index_back", backAllData)
	return
}
