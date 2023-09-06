package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"minitok/cmd/relation/global"
)

func Redis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Configs.Redis.Addr(),
		Password: global.Configs.Redis.Password, // 没有密码，默认值
		DB:       global.Configs.Redis.DB,       // 默认DB 0
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return rdb, nil
}
