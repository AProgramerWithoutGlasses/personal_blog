package dao

import (
	"staging/src/model/gorm_model"
	"time"
)

func (dao *Dao) GetPosts() (posts []gorm_model.Post, err error) {
	err = dao.db.Preload("Comments").Preload("Category").Find(&posts).Error
	return
}

func (dao *Dao) GetPrePost(createdAt time.Time) (prePost gorm_model.Post, err error) {
	err = dao.db.
		Where("Created_At < ?", createdAt).
		Order("Created_At DESC").
		First(&prePost).Error
	return
}

func (dao *Dao) GetNextPost(createdAt time.Time) (nextPost gorm_model.Post, err error) {
	err = dao.db.
		Where("Created_At > ?", createdAt).
		Order("Created_At ASC").
		First(&nextPost).Error
	return
}

func (dao *Dao) GetPostByslug(slug string) (post gorm_model.Post, err error) {
	err = dao.db.Preload("Comments").Preload("Category").First(&post, "slug = ?", slug).Error
	return
}

func (dao *Dao) GetBase() (base gorm_model.Base, err error) {
	err = dao.db.Find(&base).Error
	return
}

func (dao *Dao) GetCategories() (categories []gorm_model.Category, err error) {
	err = dao.db.Preload("Posts").Find(&categories).Error
	return
}
