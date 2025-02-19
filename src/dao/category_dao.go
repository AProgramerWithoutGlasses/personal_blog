package dao

import "blog1/src/model/gorm_model"

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
