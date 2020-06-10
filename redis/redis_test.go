package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

// https://github.com/go-redis/redis

func TestA(t *testing.T) {
	opt, err := redis.ParseURL("redis://101.133.221.239:7777/0")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	v, err := rdb.Get(context.Background(), "k1").Result()
	if err != nil {
		panic(err)
	}
	t.Log(v)
}
