package dao

import (
	"blog1/src/model/gorm_model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (dao *Dao) GetCategories() (categories []gorm_model.Category, err error) {
	err = dao.db.Preload("Posts").Find(&categories).Error
	return
}

func (dao *Dao) GetCategoryByName(name string) (category gorm_model.Category, err error) {
	err = dao.db.Preload("Posts").First(&category, "name = ?", name).Error
	return
}

func (dao *Dao) UpdateCategory(categoryId, categoryName string) (err error) {

	// 更新指定字段
	err = dao.db.Model(&gorm_model.Category{}).
		Where("id = ?", categoryId).
		Updates(gorm_model.Category{
			Name: categoryName,
		}).Error

	return
}

func (dao *Dao) DeleteCategory(categoryId int) (err error) {
	// 检查分类是否存在
	var category gorm_model.Category
	err = dao.db.First(&category, categoryId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("category not found: %w", err)
		}
		return fmt.Errorf("failed to find category: %w", err)
	}

	// 删除分类
	err = dao.db.Delete(&category).Error
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}

func (dao *Dao) CreateCategory(categoryModel gorm_model.Category) (err error) {
	var existingCategory gorm_model.Category
	// 检查是否存在未被软删除的记录
	err = dao.db.Where("name = ?", categoryModel.Name).First(&existingCategory).Error
	if err == nil {
		// 找到未被软删除的记录，返回错误
		return fmt.Errorf("分类名称 %s 已存在且未被删除", categoryModel.Name)
	}

	// 如果记录不存在或被软删除，直接创建新记录
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = dao.db.Create(&categoryModel).Error
		if err != nil {
			return err
		}
		return nil
	}

	// 其他错误
	return err
}
