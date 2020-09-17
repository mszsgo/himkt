package hmopen

import (
	"errors"
	"himkt/hm"
)

var (
	amap map[string]string
)

func getEncryptSecret(appid string) (string, error) {
	if appid == "" {
		return "", errors.New("平台appid不能为空")
	}
	secret := amap[appid]
	if secret != "" {
		return secret, nil
	}
	r, err := GetOpenApp(appid, false)
	if err != nil {
		return "", err
	}
	amap[appid] = r.DesKey
	secret = r.DesKey
	return secret, nil
}

// DES+MD5接口报文加密
func EncryptDesMd5(s, appid string) (v string, err error) {
	key, err := getEncryptSecret(appid)
	if err != nil {
		return
	}
	v, err = hm.EncryptDesMd5(s, key)
	if err == nil {
		return
	}

	amap[appid] = ""
	key, err = getEncryptSecret(appid)
	if err != nil {
		return
	}
	v, err = hm.EncryptDesMd5(s, key)
	return
}

// DES+MD5接口报文解密
func DecryptDesMd5(s, appid string) (v string, err error) {
	key, err := getEncryptSecret(appid)
	if err != nil {
		return
	}
	v, err = hm.DecryptDesMd5(s, key)
	if err == nil {
		return
	}

	amap[appid] = ""
	key, err = getEncryptSecret(appid)
	if err != nil {
		return
	}
	v, err = hm.DecryptDesMd5(s, key)
	return
}
