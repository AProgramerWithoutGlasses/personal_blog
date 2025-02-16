package service

import (
	"blog1/src/dao"
	"blog1/src/pkg/settings"
	"golang.org/x/sync/singleflight"
)

type Service struct {
	dao    *dao.Dao
	single *singleflight.Group // 合并相同的并发请求，提高性能
}

func InitService(app *settings.AppConfig) *Service {
	svc := &Service{
		dao:    dao.Init(app),
		single: new(singleflight.Group),
	}
	return svc
}
