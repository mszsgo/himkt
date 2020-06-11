package des

import (
	"encoding/hex"
	"testing"
)

func TestTRIPLE_DES_ECB_PKCS5(t *testing.T) {
	key := "111111111111111111111111"
	enVal := TRIPLE_DES_ECB_PKCS5_Encode("123", key)
	t.Log(enVal)
	deVal := TRIPLE_DES_ECB_PKCS5_Decode(enVal, key)
	t.Log(deVal)
}

func TestTRIPLE_DES_ECB_PKCS5_Unionpay(t *testing.T) {
	key := "20a140eabc4343f2c8cd62549e7a0bf420a140eabc4343f2"
	src := "13611703040"
	bytes, err := hex.DecodeString(key)
	if err != nil {
		t.Error(err)
		return
	}
	enVal := TRIPLE_DES_ECB_PKCS5_Encode(src, string(bytes))
	t.Log(enVal)
	deVal := TRIPLE_DES_ECB_PKCS5_Decode(enVal, string(bytes))
	t.Log(deVal)
}
