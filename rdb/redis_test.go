package rdb

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

	v := rdb.Get(context.TODO(), "k1").Val()
	/*if err != nil {
		panic(err)
	}*/
	t.Log(v)

	// 生成12位编号
	// 8+3+7
	// 8位日期+3位IP+7位递增

}
