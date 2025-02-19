package server

import (
	"blog1/src/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BackDelComment(c *gin.Context) {
	commentId := c.Param("comment_id")

	err := svc.BackDelCommentService(commentId)
	if err != nil {
		zap.L().Error("BackDelComment() svc.BackDelComment() err: ", zap.Error(err))
		response.Fail(c, "back_post_edit", response.ErrServer)
		return
	}

	response.SuccessWithMsg(c, "back_post_edit", "删除成功", "")

}
