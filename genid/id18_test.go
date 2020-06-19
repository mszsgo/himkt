package genid

import (
	"testing"
	"time"
)

var mn = make(map[string]int)

func TestIpmain(t *testing.T) {
	Ipmain()

	t.Log(time.Now().Nanosecond())
	for i := 1; i < 10000; i++ {
		k := string(time.Now().Nanosecond())
		if mn[k] != 0 {
			t.Log("出现重复了", i)
			break
		}
		mn[k] = i
	}
}
