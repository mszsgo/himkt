package unionpay

import (
	"context"
	"github.com/mszsgo/himkt/micro"
)

func HmUnionpayOauth2(ctx context.Context, params *HmUnionpayOauth2Params) (result *HmUnionpayOauth2Result, err error) {
	err = micro.Call(ctx, "hm.unionpay.oauth2", params, &result)
	return
}

type HmUnionpayOauth2Params struct {
	RequestId  string `json:"requestId"`
	SubmitTime string `json:"submitTime"`
	MchId      string `json:"mchId"`
	Scope      string `json:"scope" description:"可选值参考银联说明文档，如获取手机号使用：upapi_mobile"`
	Cburl      string `json:"cburl" description:"前端接收参数的URL，授权成功跳转时对URL链接结尾拼接参数“code=xxxxx”"`
}

type HmUnionpayOauth2Result struct {
	RequestId  string `json:"requestId"`
	SubmitTime string `json:"submitTime"`
	HostTime   string `json:"hostTime"`
	HostNo     string `json:"hostNo"`
	AuthUrl    string `json:"authUrl"`
}
