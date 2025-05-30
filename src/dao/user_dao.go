package dao

import (
	"blog1/src/model/gorm_model"
	"fmt"
)

func (dao *Dao) GetUserByUsername(username string) (user gorm_model.User, err error) {
	err = dao.db.Take(&user, "username = ?", username).Error
	return
}

// ExistedUser 检查User表中是否存在传入username对应的记录，若存在则返回true，不存在返回false
func (dao *Dao) ExistedUser(username string) (bool, error) {
	var count int64
	err := dao.db.Model(&gorm_model.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		fmt.Println("dao.db.Model(&gorm_model.User{}).Where().Count() err ", err)
		return false, err
	}
	return count > 0, nil
}

func (dao *Dao) InsertUser(registerModel gorm_model.User) (err error) {
	return dao.db.Create(&registerModel).Error
}

func (dao *Dao) EditUser(user gorm_model.User) (err error) {
	err = dao.db.Model(gorm_model.User{}).
		Where("id = ?", user.ID).
		Select("Name", "Age", "Gender", "Permission"). // 显式指定要更新的字段
		Updates(&user).Error
	return
}
