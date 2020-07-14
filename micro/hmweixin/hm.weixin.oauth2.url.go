package hmweixin

import (
	"context"
	"github.com/mszsgo/himkt/micro"
)

func HmWeixinOauth2Url(ctx context.Context, params *HmWeixinOauth2UrlParams) (result *HmWeixinOauth2UrlResult, err error) {
	err = micro.Call(ctx, "hm.weixin.oauth2.url", params, &result)
	return
}

type HmWeixinOauth2UrlParams struct {
	MchId       string `json:"mchId" description:"根据mchId查询公众号配置信息"`
	Scope       string `json:"scope" description:"微信授权范围，根据微信文档定义传值"`
	RedirectUri string `json:"redirectUri" description:"获取到微信网页授权openid后，跳转的网页地址，拼接code参数，从hm-weixin服务查询用户信息"`
}

type HmWeixinOauth2UrlResult struct {
	Url string `json:"url"`
}
