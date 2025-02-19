package dao

import (
	"blog1/src/model/gorm_model"
	"github.com/pkg/errors"
	"time"
)

func (dao *Dao) GetPosts() (posts []gorm_model.Post, err error) {
	err = dao.db.Preload("Comments").Preload("Category").Find(&posts).Error
	return
}

func (dao *Dao) GetPostsByCategory(categoryName string) (posts []gorm_model.Post, err error) {

	// 使用 Joins 和 Where 查询
	err = dao.db.Model(&gorm_model.Post{}).Joins("JOIN categories ON posts.category_id = categories.id").
		Where("categories.name = ?", categoryName).
		Find(&posts).Error

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

func (dao *Dao) GetPostBySlug(slug string) (post gorm_model.Post, err error) {
	err = dao.db.Preload("Comments").Preload("Category").First(&post, "slug = ?", slug).Error
	return
}

func (dao *Dao) DeletePostBySlug(slug string) (err error) {
	return dao.db.Delete(&gorm_model.Post{}, "slug = ?", slug).Error
}

func (dao *Dao) GetBase() (base gorm_model.Base, err error) {
	err = dao.db.Find(&base).Error
	return
}

func (dao *Dao) GetComments() (comments []gorm_model.Comment, err error) {
	err = dao.db.Preload("User").Preload("Post").Find(&comments).Error
	return
}

func (dao *Dao) GetUsers() (users []gorm_model.User, err error) {
	err = dao.db.Preload("Comments").Find(&users).Error
	return
}

func (dao *Dao) InsertPost(post gorm_model.Post) (err error) {
	err = dao.db.Create(&post).Error
	return
}

func (dao *Dao) UpdatePost(post gorm_model.Post, slug string) (err error) {
	// 确保 slug 不为空
	if slug == "" {
		return errors.New("slug 不能为空")
	}

	// 更新指定字段
	err = dao.db.Model(&gorm_model.Post{}).
		Where("slug = ?", slug).
		Updates(gorm_model.Post{
			Title:      post.Title,
			Content:    post.Content,
			CoverImg:   post.CoverImg,
			Summary:    post.Summary,
			CategoryID: post.CategoryID,
		}).Error

	return
}
