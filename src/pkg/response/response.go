package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, name string, data any) {
	c.HTML(http.StatusOK, name, response{
		Code: 0,
		Msg:  "",
		Data: data,
	})
	c.Abort()
}

func Fail(c *gin.Context, name string, msg string) {
	c.HTML(http.StatusInternalServerError, name, response{
		Code: 1,
		Msg:  msg,
		Data: nil,
	})
	c.Abort()
}
