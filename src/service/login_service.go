package service

import (
	"blog1/src/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (s *Service) LoginService(loginModel model.LoginReqModel) (myPermission string, err error) {
	user, err := s.dao.GetUserByUsername(loginModel.Username)
	if err != nil {
		fmt.Println("s.dao.GetUserByUsername(loginModel.Username) err: ", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = model.ErrUserNotExist
		}
		return
	}

	if loginModel.Password != user.Password {
		err = model.ErrInvalidPassword
	}

	myPermission = user.Permission

	return
}
