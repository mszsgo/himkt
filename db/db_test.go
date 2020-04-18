package db

import (
	"github.com/mszsgo/himkt/env"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Log(Connect(env.HM_MONGO_CONNECTION_STRING))
}
