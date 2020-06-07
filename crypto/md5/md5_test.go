package md5

import "testing"

func TestEncrypt(t *testing.T) {
	t.Log(Encrypt("123"))
}
