package service

import (
	"blog1/src/model"
	"blog1/src/pkg/response"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

	// 校验密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginModel.Password))
	if err != nil {
		// 密码错误
		fmt.Println("LoginService() bcrypt.CompareHashAndPassword() err: ", err.Error())
		err = errors.New(response.ErrInvalidPassword.Msg())
		return
	}

	myPermission = user.Permission

	return
}
