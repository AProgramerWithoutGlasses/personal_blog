package middleware

import (
	"github.com/gin-gonic/gin"
	"staging/src/pkg/jwt"
	"staging/src/pkg/response"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件，用于获得、验证token，并解析出自定义信息放入上下文。
// 1. Authorization字段值是否为空； 2.开头是否为Bearer； 3.token解析是否错误； 4. 解析得到token的附带自定义信息填入上下文。
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定

		// 从请求头中获取到token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(c, "err.html", "未提供认证令牌token")
		}

		// 按空格分割 检查开头是否为Bearer
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(c, "err.html", "token格式错误")
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			response.Fail(c, "err.html", "token解析错误")
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("permission", mc.Permission)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
