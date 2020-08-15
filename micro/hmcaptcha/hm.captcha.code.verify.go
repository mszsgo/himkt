package hmcaptcha

import (
	"context"
	"himkt/micro"
)

func HmCaptchaCodeVerify(ctx context.Context, params *HmCaptchaCodeVerifyParams) (r *HmCaptchaCodeVerifyResult, err error) {
	err = micro.Call(ctx, "hm.captcha.code.verify", params, &r)
	return
}

type HmCaptchaCodeVerifyParams struct {
	CaptchaId string `json:"captchaId" description:"验证码编号"`
	Value     string `json:"value" description:"客户端输入的验证码值"`
}

type HmCaptchaCodeVerifyResult struct {
}
