package dao

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wushengyouya/project-project/config"
)

var RC *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(config.AppConf.InitRedisOptions())
	RC = &RedisCache{rdb: rdb}
}

func (r *RedisCache) Put(key, value string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return r.rdb.Set(ctx, key, value, expire).Err()

}

func (r *RedisCache) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return r.rdb.Get(ctx, key).Result()
}
