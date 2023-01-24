package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func StartRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var ctx = context.Background()
	err := rdb.Set(ctx, "message", "ex", 0).Err()
	rdb.Set(ctx, "zalupa", "wedw", 90000)
	if err != nil {
		panic(err)
	}
}
