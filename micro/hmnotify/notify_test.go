package hmnotify

import (
	"himkt/genid"
	"testing"
)

func TestRequest(t *testing.T) {
	bm := make(map[string]string)
	bm["requestId"] = genid.UUID()
	err := Request("https://msd.himkt.cn/work/trade/testing/async", "88888888", "tt.args.method", bm)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("success....................")
}
