package hmsms

import (
	"context"
	"himkt/micro"
)

// 用于发送验证码类短信
func HmSmsCodeVerify(ctx context.Context, params *HmSmsCodeVerifyParams) (result *HmSmsCodeVerifyResult, err error) {
	err = micro.Call(ctx, "hm.sms.code.verify", params, &result)
	return
}

type HmSmsCodeVerifyParams struct {
	RequestId  string `json:"requestId"`
	SubmitTime string `json:"submitTime"`
	OrgId      string `json:"orgId"`
	Mobile     string `json:"mobile"`
	CodeId     string `json:"codeId"`
	Code       string `json:"code"`
}

type HmSmsCodeVerifyResult struct {
	RequestId string `json:"requestId"`
}
