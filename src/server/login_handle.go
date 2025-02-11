package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"staging/src/model"
	"staging/src/pkg/jwt"
	"staging/src/pkg/response"
)

// loginFrontHandler 用于登录前台
func loginFrontHandler(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(1111)
		// 接收参数
		var loginReqModel = model.LoginReqModel{}
		err := c.ShouldBindJSON(&loginReqModel)
		if err != nil {
			fmt.Println("c.ShouldBindJSON(&loginReqModel) err:", err)
			response.Fail(c, name, err.Error())
			zap.L().Error("c.ShouldBindJSON(&loginModel) err: ", zap.Error(err))
			return
		}

		// 业务
		myPermission, err := svc.LoginService(loginReqModel)
		if err != nil {
			fmt.Println("svc.LoginService(&loginReqModel) err:", err)
			response.Fail(c, name, err.Error())
			zap.L().Error("svc.LoginService err: ", zap.Error(err))
			return
		}

		// 设置token
		token, err := jwt.GenerateToken(loginReqModel.Username, myPermission)
		if err != nil {
			fmt.Println("jwt.GenerateToken(&loginReqModel) err:", err)
			response.Fail(c, name, err.Error())
			zap.L().Error("jwt.GenerateToken err: ", zap.Error(err))
			return
		}
		fmt.Println(2222)

		// 响应
		loginRes := model.LoginResModel{Token: token}
		response.Success(c, name, loginRes)
		fmt.Println(33333)

	}
}

// loginFrontHandler 用于登录前台
func loginFrontView(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		response.Success(c, name, "登录成功")
	}
}

// loginBackHandler 用于登录后台
func loginBackHandler(c *gin.Context) {}
