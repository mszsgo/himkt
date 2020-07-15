package hmsms

import (
	"context"
	"github.com/mszsgo/himkt/micro"
)

// 用于发送验证码类短信
func HmSmsCodeSend(ctx context.Context, params *HmSmsCodeSendParams) (result *HmSmsCodeSendResult, err error) {
	err = micro.Call(ctx, "hm.sms.code.send", params, &result)
	return
}

type HmSmsCodeSendParams struct {
	RequestId  string `json:"requestId"`
	SubmitTime string `json:"submitTime"`
	OrgId      string `json:"orgId"`
	Mobile     string `json:"mobile"`
	TemplateId string `json:"templateId"` // 模板中只能存在一个${code}参数
	CodeLen    int    `json:"codeLen"`    // 验证码长度，默认6
}

type HmSmsCodeSendResult struct {
	RequestId string `json:"requestId"`
	TradeId   string `json:"tradeId"`
	CodeId    string `json:"codeId"`
}
