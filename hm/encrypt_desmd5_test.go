package hm

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestEncryptDesMd5(t *testing.T) {
	key := "88888888"
	s := `{"appid":"88888888"}`
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
	t.Log(url.QueryEscape(v))

	// 解密
	v, e = DecryptDesMd5(v, key)
	if e != nil {
		t.Error(e.Error())
		return
	}
	t.Log(v)
}

type S1 struct {
	Aa string `json:"aa"`
}

type S2 struct {
	Bb string `json:"bb"`
	S1
}

func TestJ(t *testing.T) {
	v := &S2{Bb: "123", S1: S1{Aa: "321"}}
	b, _ := json.Marshal(v)
	t.Log(string(b))

	var s2 *S2
	jv := []byte(`{"bb":"123","aa":"321"}`)
	json.Unmarshal(jv, &s2)
	s2b, _ := json.Marshal(s2)

	s2b = []byte("{}")
	suc := `{"errno":"00000","error":"ok"}`
	if len(s2b) > 2 {
		s2b = append([]byte(suc[0:len(suc)-1]+","), s2b[1:]...)
	} else {
		s2b = []byte(suc)
	}
	t.Log(string(s2b))
}
