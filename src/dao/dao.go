package dao

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"staging/src/pkg/settings"
)

type Dao struct {
	db  *gorm.DB
	rdb *redis.Client
}

func Init(app *settings.AppConfig) *Dao {
	dao := &Dao{
		db:  initDB(app.MySQLConfig),
		rdb: initRDB(app.RedisConfig),
	}
	return dao
}
