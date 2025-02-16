package dao

import (
	"blog1/src/model/gorm_model"
	"blog1/src/pkg/settings"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB(m *settings.MySQLConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		zap.L().Error("gorm init failed", zap.Error(err)) // 修正日志输出
	}

	// 自动迁移
	err = db.AutoMigrate(&gorm_model.User{}, &gorm_model.Post{}, &gorm_model.Category{}, &gorm_model.Comment{}, &gorm_model.Base{})
	if err != nil {
		zap.L().Error("gorm migrate failed", zap.Error(err))
		return nil
	}

	return db
}
