package datasource

import (
	"github.com/go-redis/redis"
	"gomall/pkg/settings"
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

func GetReids(domain string) *redis.Client {
	return redisMaps[domain]
}
