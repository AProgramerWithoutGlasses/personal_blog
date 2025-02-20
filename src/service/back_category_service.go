package service

import (
	"blog1/src/model"
	"blog1/src/model/gorm_model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func (s *Service) BackCategoriesService() (backCategoriesModel model.BackCategoriesModel, err error) {
	categories, err := s.dao.GetCategories()
	if err != nil {
		fmt.Println("CategoryService() s.dao.GetCategories() err:", err)
	}

	base, err := s.dao.GetBase()
	if err != nil {
		fmt.Println("s.dao.GetBase() err: ", err)
		return
	}

	backCategoriesModel = model.BackCategoriesModel{
		Categories:  categories,
		Base:        base,
		CurrentPage: "categories",
	}

	return
}

func (s *Service) BackEditCategoryService(categoryId string, name string) (err error) {
	err = s.dao.UpdateCategory(categoryId, name)
	if err != nil {
		fmt.Println("BackEditCategoryService() err:", err)
	}
	return
}

func (s *Service) BackDelCategoryService(categoryId string) (err error) {
	// String 转为 Int
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		fmt.Println("BackDelCategoryService() err:", err)
		return
	}

	// 调用dao层进行删除
	err = s.dao.DeleteCategory(categoryIdInt)
	if err != nil {
		fmt.Println("BackEditCategoryService() err:", err)
		return
	}

	return
}

func (s *Service) BackNewCategoryService(backNewCategoriesModel model.BackNewCategoriesModel) (err error) {
	categoryModel := gorm_model.Category{Name: backNewCategoriesModel.Name}

	// 调用dao层进行增加
	err = s.dao.CreateCategory(categoryModel)
	if err != nil {
		fmt.Println("s.dao.CreateCategory() err:", err)
		return
	}

	return
}

func (s *Service) BackUsersViewService() (backUsersModel model.BackUsersModel, err error) {
	users, err := s.dao.GetUsers()
	if err != nil {
		fmt.Println("s.dao.GetUsers() err:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
	}

	base, err := s.dao.GetBase()
	if err != nil {
		fmt.Println("s.dao.GetBase()", err)
		return
	}

	backUsersModel = model.BackUsersModel{
		Users:       users,
		Base:        base,
		CurrentPage: "users",
	}

	return
}

func (s *Service) BackDelUsersService() (backUsersModel model.BackUsersModel, err error) {
	s.dao.user
	return
}
