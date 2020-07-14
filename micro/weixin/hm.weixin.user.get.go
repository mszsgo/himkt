package src

import (
	"context"
	"github.com/mszsgo/himkt/micro"
)

func HmWeixinUserGet(ctx context.Context, params *HmWeixinUserGetParams) (result *HmWeixinUserGetResult, err error) {
	err = micro.Call(ctx, "hm.weixin.user.get", params, &result)
	return
}

type HmWeixinUserGetParams struct {
	Code string `json:"code"`
}

type HmWeixinUserGetResult struct {
	MchId      string `json:"mchId"`
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        int64  `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
}
