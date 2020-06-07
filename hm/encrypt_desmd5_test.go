package hm

import "testing"

func TestEncryptDesMd5(t *testing.T) {
	key := "12345678"
	s := "123中文"
	/*
		密文 e2a19818214e2c1523208747dd872fa9xmF4oTvCx2mrAp8erL7oHg==
		明文 123中文
	*/

	//加密
	v, e := EncryptDesMd5(s, key)
	if e != nil {
		t.Error(e.Error())
		return
	}
	t.Log(v)

	// 解密
	v, e = DecryptDesMd5(v, key)
	if e != nil {
		t.Error(e.Error())
		return
	}
	t.Log(v)
}
