package des

import "testing"

func TestDES_ECB_PKCS5(t *testing.T) {
	key := "20a140eabc4343f2c8cd62549e7a0bf420a140eabc4343f2"
	enVal := DES_ECB_PKCS5_Encode("123", key)
	t.Log(enVal)
	deVal := DES_ECB_PKCS5_Decode(enVal, key)
	t.Log(deVal)
}
