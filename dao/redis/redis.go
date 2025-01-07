package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Init() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return
}

func Close() {
	_ = client.Close()
}
