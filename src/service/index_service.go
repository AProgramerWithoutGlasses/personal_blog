package service

import (
	"blog1/src/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (s *Service) IndexService() (res model.IndexModel, err error) {
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

	res = model.IndexModel{
		Posts:      posts,
		Base:       base,
		Categories: categories,
	}

	return
}

func (s *Service) PostService(slug string) (postModel model.PostModel, err error) {
	post, err := s.dao.GetPostBySlug(slug)
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

	categories, err := s.dao.GetCategories()
	if err != nil {
		fmt.Println("get categories err:", err)
		return
	}

	postModel = model.PostModel{
		Post:       post,
		PrePost:    prePost,
		NextPost:   nextPost,
		Base:       base,
		Categories: categories,
	}

	return
}

func (s *Service) CategoryService(categoryName string) (categoryModel model.CategoryModel, err error) {
	currentCategory, err := s.dao.GetCategoryByName(categoryName)
	if err != nil {
		fmt.Println("get category err:", err)
		return
	}

	currentPosts, err := s.dao.GetPostsByCategory(categoryName)
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

	categoryModel = model.CategoryModel{
		CurrentCategory: currentCategory,
		CurrentPosts:    currentPosts,
		Base:            base,
		Categories:      categories,
		PostsCount:      len(currentPosts),
	}

	return
}
