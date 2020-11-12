package service

import (
	"context"
	"time"

	"github.com/brmyfun/brmy-go/config"
)

var ctx = context.Background()

// RedisSet redis设置数据
func RedisSet(key string, value interface{}, expiration time.Duration) (string, error) {
	return config.Rdb.Set(ctx, key, value, expiration).Result()
}

// RedisGet redis根据key值获取数据
func RedisGet(key string) (interface{}, error) {
	return config.Rdb.Get(ctx, key).Result()
}
