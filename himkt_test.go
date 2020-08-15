package himkt

import (
	"himkt/cfg"
	"himkt/db"
	"himkt/env"
	"himkt/hm"
	"testing"
)

func TestC(t *testing.T) {
	t.Log(env.HM_ENV)
	t.Log(cfg.Serve())
	t.Log(db.M{})
	t.Log(hm.Track{})
}
