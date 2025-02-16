package service

import "blog1/src/model"

func (s *Service) Test() (data model.Test, err error) {
	data, err = s.dao.Test()
	return
}
