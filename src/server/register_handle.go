package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"staging/src/model"
	"staging/src/pkg/response"
)

func registerHandler(name string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 接收参数
		var registerModel = model.RegisterReqModel{}
		err := c.ShouldBindJSON(&registerModel)
		if err != nil {
			fmt.Println("c.ShouldBindJSON(&registerModel)", err)
			zap.L().Error("c.ShouldBindJSON(&registerModel)", zap.Error(err))
			response.Fail(c, name, "无效的参数")
			return
		}

		// 业务
		err = svc.RegisterService(&registerModel)
		if err != nil {
			fmt.Println("svc.RegisterService(&registerModel)", err)
			zap.L().Error("svc.RegisterService(&registerModel)", zap.Error(err))
			response.Fail(c, name, err.Error())
			return
		}

		// 响应
		response.Success(c, name, registerModel.Username+"注册成功")

		return
	}
}
