package redis

import "github.com/redis/go-redis/v9"

var r *redis.Client

func Init() {
	r = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
}
