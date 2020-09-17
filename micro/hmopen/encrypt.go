package hmopen

import (
	"himkt/hm"
)

// DES+MD5接口报文加密
func EncryptDesMd5(s, appid string) (v string, err error) {
	key, err := GetDesKey(appid)
	if err != nil {
		return
	}
	v, err = hm.EncryptDesMd5(s, key)
	if err == nil {
		return
	}

	key, err = GetDesKeyN(appid)
	if err != nil {
		return
	}
	v, err = hm.EncryptDesMd5(s, key)
	return
}

// DES+MD5接口报文解密
func DecryptDesMd5(s, appid string) (v string, err error) {
	key, err := GetDesKey(appid)
	if err != nil {
		return
	}
	v, err = hm.DecryptDesMd5(s, key)
	if err == nil {
		return
	}

	key, err = GetDesKeyN(appid)
	if err != nil {
		return
	}
	v, err = hm.DecryptDesMd5(s, key)
	return
}
