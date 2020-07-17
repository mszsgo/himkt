package hmopen

import (
	"context"
	"github.com/mszsgo/himkt/micro"
	"sync"
)

var (
	apps sync.Map
)

func GetOpenApp(appid string, readCache bool) (r *HmOpenAppCryptoResult, err error) {
	if readCache {
		value, ok := apps.Load(appid)
		if ok {
			r = value.(*HmOpenAppCryptoResult)
			return
		}
	}
	r, err = HmOpenAppCrypto(nil, &HmOpenAppCryptoParams{Appid: appid})
	if err != nil {
		return
	}
	apps.Store(appid, r)
	return
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
