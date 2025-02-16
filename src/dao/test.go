package dao

import (
	"blog1/src/model"
	"context"
)

func (dao *Dao) Test() (test model.Test, err error) {
	err = dao.db.Find(&test).Error
	dao.rdb.Get(context.Background(), "test")
	return
}
