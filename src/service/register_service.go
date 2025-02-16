package service

import (
	"blog1/src/model"
	"blog1/src/model/gorm_model"
	"blog1/src/pkg/response"
	"errors"
	"fmt"
	"strconv"
)

func (s *Service) RegisterService(registerModel *model.RegisterReqModel) (err error) {
	// 先判断该用户是否存在
	existed, err := s.dao.ExistedUser(registerModel.Username)
	if err != nil {
		fmt.Println("dao.ExistedUser err:", err)
		return
	}

	if existed {
		// 若该用户已存在，则返回，不再注册。
		err = errors.New(response.ErrUserExisted.Msg())
	} else {
		// 若该用户不存在，则进行注册
		user := gorm_model.User{
			Username: registerModel.Username,
			Password: registerModel.Password,
			Name:     registerModel.Name,
			Gender:   registerModel.Gender,
			Age:      strconv.Itoa(registerModel.Age),
		}

		err = s.dao.InsertUser(user)
		if err != nil {
			fmt.Println("dao.InsertUser err:", err)
			return
		}
	}

	return
}
