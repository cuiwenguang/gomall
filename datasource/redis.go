package datasource

import (
	"github.com/go-redis/redis"
	"gomall/pkg/settings"
	"gomall/pkg/web"
)

var redisMaps map[string]*redis.Client

func initRedis() {
	redisMaps = make(map[string]*redis.Client)
	for k, v := range settings.AppConfig.Redis {
		client := redis.NewClient(&redis.Options{
			Addr:     v.Addr,
			Password: v.Password,
			DB:       v.DB,
		})
		redisMaps[k] = client
	}
}

func GetReids(ctx *web.RequestContext) *redis.Client {
	domain := GetDomain(ctx.Host)
	return redisMaps[domain]
}
