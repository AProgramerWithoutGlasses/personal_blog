package service

import (
	"blog1/src/model"
	"fmt"
)

func (s *Service) BackCategoriesService() (backCategoriesModel model.BackCategoriesModel, err error) {
	categories, err := s.dao.GetCategories()
	if err != nil {
		fmt.Println("CategoryService() s.dao.GetCategories() err:", err)
	}

	base, err := s.dao.GetBase()
	if err != nil {
		fmt.Println("get base err:", err)
		return
	}

	backCategoriesModel = model.BackCategoriesModel{
		Categories: categories,
		Base:       base,
	}

	return
}

func (s *Service) BackEditCategoryService(categoryId, categoryName string) (backCategoriesModel model.BackCategoriesModel, err error) {
	err = s.dao.UpdateCategory(categoryId, categoryName)
	if err != nil {
		fmt.Println("BackEditCategoryService() err:", err)
	}
	return
}
