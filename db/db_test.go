package db

import (
	"github.com/mszsgo/himkt"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Log(Connect(himkt.HM_MONGO_CONNECTION_STRING))
}
