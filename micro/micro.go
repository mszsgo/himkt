package micro

import (
	"context"
	"github.com/mszsgo/himkt/hm"
)

func Call(ctx context.Context, method string, i interface{}, o interface{}) error {
	return hm.Call(ctx, method, i, o)
}
