package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"web-app/settings"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

func Init(conf *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.Dbname,
		Protocol: conf.PoolSize,
	})
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return
}

func Close() {
	_ = client.Close()
}
