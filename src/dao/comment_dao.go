package dao

import "blog1/src/model/gorm_model"

func (dao *Dao) DelCommentById(id string) (err error) {
	return dao.db.Delete(&gorm_model.Comment{}, "id = ?", id).Error
}
