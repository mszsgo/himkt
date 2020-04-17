package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// 使用示例： himkt.M{}.Str("name",name).Str("k",v)
// 自定义Map，忽略空值插入
type M bson.M

func (f M) M(k string, v M) M {
	if len(v) > 0 {
		f[k] = v
	}
	return f
}

func (f M) Str(k string, v string) M {
	if v != "" {
		f[k] = v
	}
	return f
}

func (f M) Int(k string, v int) M {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f M) Int32(k string, v int32) M {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f M) Int64(k string, v int64) M {
	if v != 0 {
		f[k] = v
	}
	return f
}

func (f M) Time(k string, v time.Time) M {
	if !v.IsZero() {
		f[k] = v
	}
	return f
}
