package unionpay

import (
	"context"
	"github.com/mszsgo/himkt/micro"
)

func HmUnionpayUserGet(ctx context.Context, params *HmUnionpayUserGetParams) (result *HmUnionpayUserGetResult, err error) {
	err = micro.Call(ctx, "hm.unionpay.user.get", params, &result)
	return
}

type HmUnionpayUserGetParams struct {
	Code string `json:"code"`
}

type HmUnionpayUserGetResult struct {
	MchId  string `json:"mchId"`
	OpenId string `json:"openId"`
	Mobile string `json:"mobile"`
	AppId  string `json:"appId"`
}
