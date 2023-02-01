package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func StartRedis() {
	// todo don't panic, return err and rdb. Rename redis.StartRedis() -> redis.connect()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // todo get values for redis.Options from yaml config
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	var ctx = context.Background()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}
