package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"staging/src/model"
)

func (s *Service) IndexService() (res model.IndexModel, err error) {
	posts, err := s.dao.GetPosts()
	if err != nil {
		fmt.Println("get posts err:", err)
		return
	}
	fmt.Println("posts:", posts)

	base, err := s.dao.GetBase()
	if err != nil {
		fmt.Println("get base err:", err)
		return
	}
	fmt.Println("base:", base)

	categories, err := s.dao.GetCategories()
	if err != nil {
		fmt.Println("get categories err:", err)
		return
	}
	fmt.Println("categories:", categories)

	res = model.IndexModel{
		Posts:      posts,
		Base:       base,
		Categories: categories,
	}

	return
}

func (s *Service) PostService(slug string) (postModel model.PostModel, err error) {
	post, err := s.dao.GetPostByslug(slug)
	if err != nil {
		fmt.Println("get post err:", err)
		return
	}

	prePost, err := s.dao.GetPrePost(post.CreatedAt)
	if err != nil {
		fmt.Println("get prePost err:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
	}

	nextPost, err := s.dao.GetNextPost(post.CreatedAt)
	if err != nil {
		fmt.Println("get nextPost err:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
	}

	base, err := s.dao.GetBase()
	if err != nil {
		fmt.Println("get base err:", err)
		return
	}
	fmt.Println("base:", base)

	categories, err := s.dao.GetCategories()
	if err != nil {
		fmt.Println("get categories err:", err)
		return
	}
	fmt.Println("categories:", categories)

	postModel = model.PostModel{
		Post:       post,
		PrePost:    prePost,
		NextPost:   nextPost,
		Base:       base,
		Categories: categories,
	}

	fmt.Println("postModel:", postModel)

	return
}
