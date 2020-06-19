package rdb

import "testing"

func TestKeyPre_Inc(t *testing.T) {
	gid := KeyPre("inc:")

	t.Log(gid.Inc("tt"))
	t.Log(gid.Inc("tt"))
}
