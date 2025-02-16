package service

import (
	"fmt"
	"staging/src/model"
)

func (s *Service) BackAllDateService() (allBackDate model.BackAllDataModel, err error) {
	posts, err := s.dao.GetPosts()
	if err != nil {
		fmt.Println("get posts err:", err)
		return
	}

	base, err := s.dao.GetBase()
	if err != nil {
		fmt.Println("get base err:", err)
		return
	}

	categories, err := s.dao.GetCategories()
	if err != nil {
		fmt.Println("get categories err:", err)
		return
	}

	comments, err := s.dao.GetComments()
	if err != nil {
		fmt.Println("get comments err:", err)
		return
	}

	users, err := s.dao.GetUsers()
	if err != nil {
		fmt.Println("get users err:", err)
		return
	}

	allBackDate = model.BackAllDataModel{
		Posts:      posts,
		Base:       base,
		Categories: categories,
		Comments:   comments,
		Users:      users,
	}
	return
}
