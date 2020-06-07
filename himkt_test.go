package himkt

import (
	"github.com/mszsgo/himkt/cfg"
	"github.com/mszsgo/himkt/db"
	"github.com/mszsgo/himkt/env"
	"github.com/mszsgo/himkt/hm"
	"testing"
)

func TestC(t *testing.T) {
	t.Log(env.HM_ENV)
	t.Log(cfg.Serve())
	t.Log(db.M{})
	t.Log(hm.Track{})
}
