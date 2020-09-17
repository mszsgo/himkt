package hmopen

import (
	"context"
	"errors"
	"himkt/micro"
	"sync"
)

var (
	apps sync.Map
)

func GetDesKey(appid string) (desKey string, err error) {
	if appid == "" {
		return "", errors.New("平台appid不能为空")
	}
	val, ok := apps.Load(appid)
	if ok {
		desKey = val.(string)
		return
	}
	desKey, err = GetDesKeyN(appid)
	if err != nil {
		return
	}
	return
}

func GetDesKeyN(appid string) (desKey string, err error) {
	r, err := HmOpenAppCrypto(nil, &HmOpenAppCryptoParams{Appid: appid})
	if err != nil {
		return
	}
	apps.Store(appid, r.DesKey)
	return r.DesKey, err
}

func HmOpenAppCrypto(ctx context.Context, params *HmOpenAppCryptoParams) (result *HmOpenAppCryptoResult, err error) {
	err = micro.Call(ctx, "hm.open.app.crypto", params, &result)
	return
}

type HmOpenAppCryptoParams struct {
	RequestId string `json:"requestId"`
	Appid     string `json:"appid"`
}

type HmOpenAppCryptoResult struct {
	RequestId string `json:"requestId"`
	Appid     string `json:"appid"`
	Name      string `json:"name"`
	DesKey    string `json:"desKey"`
	AesKey    string `json:"aesKey"`
	RsaPubKey string `json:"rsaPubKey"`
	RsaPriKey string `json:"rsaPriKey"`
}
