package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mszsgo/himkt/cfg"
	"time"
)

func New() *redis.Client {
	opt, err := redis.ParseURL(cfg.Redis().Url)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}

type KeyPre string

var _rdb *redis.Client

func (k KeyPre) DB() *redis.Client {
	if _rdb != nil {
		return _rdb
	}
	_rdb = New()
	return _rdb
}

func (k KeyPre) Set(key, value string, second int64) error {
	return k.DB().Set(context.TODO(), string(k)+key, value, time.Duration(second)*time.Second).Err()
}

func (k KeyPre) Get(key string) string {
	return k.DB().Get(context.TODO(), string(k)+key).Val()
}
