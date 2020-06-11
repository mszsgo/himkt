package des

import (
	"testing"
)

func TestEncryptCBC(t *testing.T) {
	v, _ := EncryptCBC([]byte("123"), []byte("12345678"))
	t.Log(v)
	v, err := DecryptCBC("6u1n1q7jQ5A=", []byte("12345678"))
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(v)
}
