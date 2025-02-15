package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, name string, data any) {
	c.HTML(http.StatusOK, name, response{
		Code: SuccessCode,
		Msg:  "",
		Data: data,
	})
	c.Abort()
}

func SuccessWithMsg(c *gin.Context, name string, msg string, data any) {
	response1 := response{
		Code: SuccessCode,
		Msg:  msg,
		Data: data,
	}
	c.HTML(http.StatusOK, name, response1)
	fmt.Printf("%#v\n", response1)
	c.Abort()
}

func Fail(c *gin.Context, name string, code Code) {
	response1 := response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.HTML(http.StatusInternalServerError, name, response1)
	fmt.Printf("%#v\n", response1)
	c.Abort()
}

// FailWithMsg 响应错误信息(自行传入msg参数)
func FailWithMsg(c *gin.Context, name string, code Code, msg string) {
	c.HTML(http.StatusInternalServerError, name, response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
	c.Abort()
}
