package service

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gomall/pkg/web"
	"gomall/storage"
)

// Serivce 基础结构
type Service struct {
	*web.RequestContext
	*gorm.DB
	*redis.Client
}

func InitService(ctx *web.RequestContext) Service {
	srv := Service{
		ctx,
		storage.GetDB(ctx),
		storage.GetReids(ctx),
	}
	return srv
}
