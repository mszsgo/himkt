package hmorg

import (
	"context"
	"himkt/micro"
	"himkt/time/t14"
)

// 查询机构信息
func HmOrgInfo(ctx context.Context, params *HmOrgInfoParams) (result *HmOrgInfoResult, err error) {
	params.SubmitTime = t14.NowF14()
	err = micro.Call(ctx, "hm.org.info", params, &result)
	return
}

type HmOrgInfoParams struct {
	RequestId  string `json:"requestId"`
	SubmitTime string `json:"submitTime"`
	OrgId      string `json:"orgId"`
}

type HmOrgInfoResult struct {
	RequestId     string `json:"requestId"`
	SubmitTime    string `json:"submitTime"`
	HostTime      string `json:"hostTime"`
	OrgId         string `json:"orgId"`
	OrgName       string `json:"orgName"`       // 机构名称
	Brand         string `json:"brand"`         // 品牌名称
	Logo          string `json:"logo"`          // 品牌LOGO 图片URL
	WeixinMchId   string `json:"weixinMchId"`   // 微信接入商户编号
	UnionpayMchId string `json:"unionpayMchId"` // 云闪付接入商户编号
	AlipayMchId   string `json:"alipayMchId"`   // 支付宝接入商户编号
}
