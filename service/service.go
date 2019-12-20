package service

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gomall/datasource"
	"gomall/pkg/web"
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
		datasource.GetDB(ctx),
		datasource.GetReids(ctx),
	}
	return srv
}
