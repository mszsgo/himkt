package hmsms

import (
	"context"
	"github.com/mszsgo/himkt/micro"
)

// 用于发送验证码类短信
func HmSmsMessageSend(ctx context.Context, params *HmSmsMessageSendParams) (result *HmSmsMessageSendResult, err error) {
	err = micro.Call(ctx, "hm.sms.message.send", params, &result)
	return
}

type HmSmsMessageSendParams struct {
	RequestId      string `json:"requestId"`
	SubmitTime     string `json:"submitTime"`
	OrgId          string `json:"orgId"`
	Mobile         string `json:"mobile"`
	TemplateId     string `json:"templateId"`
	TemplateParams string `json:"templateParams"` //模板参数 ${key}
}

type HmSmsMessageSendResult struct {
	RequestId string `json:"requestId"`
	TradeId   string `json:"tradeId"`
}
