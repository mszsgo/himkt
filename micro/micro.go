package micro

import (
	"context"
	"himkt/hm"
)

func Call(ctx context.Context, method string, i interface{}, o interface{}) error {
	return hm.Call(ctx, method, i, o)
}
