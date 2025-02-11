package dao

import (
	"fmt"
	"staging/src/model"
	"staging/src/model/gorm_model"
)

func (dao *Dao) GetUserByUsername(username string) (user gorm_model.User, err error) {
	err = dao.db.Take(&user, "username=?", username).Error
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

func (dao *Dao) InsertUser(registerModel *model.RegisterReqModel) (err error) {
	return dao.db.Create(&registerModel).Error
}
