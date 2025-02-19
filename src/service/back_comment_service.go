package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (s *Service) BackDelCommentService(commentId string) (err error) {
	err = s.dao.DelCommentById(commentId)
	if err != nil {
		fmt.Println("s.dao.DelCommentById() err:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
	}
	return
}
