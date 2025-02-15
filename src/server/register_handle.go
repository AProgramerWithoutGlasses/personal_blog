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
		err := c.ShouldBind(&registerModel)
		if err != nil {
			fmt.Println("c.ShouldBindJSON(&registerModel)", err)
			zap.L().Error("c.ShouldBindJSON(&registerModel)", zap.Error(err))
			response.Fail(c, name, response.ErrParamBind)
			return
		}

		// 业务
		err = svc.RegisterService(&registerModel)
		if err != nil {
			fmt.Println("svc.RegisterService(&registerModel)", err)
			zap.L().Error("svc.RegisterService(&registerModel)", zap.Error(err))
			response.Fail(c, name, response.ErrServer)
			return
		}

		// 响应
		response.SuccessWithMsg(c, name, registerModel.Username+"注册成功", "")

		return
	}
}

func registerView(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		response.Success(c, name, "")
	}

}
