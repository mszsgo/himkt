package hmmember

import (
	"context"
	"github.com/mszsgo/himkt/hm"
)

func HmMemberUserInfo(ctx context.Context, params *HmMemberUserInfoParams) (result *HmMemberUserInfoResult, err error) {
	err = hm.Call(ctx, "hm.member.user.info", params, &result)
	return
}

type HmMemberUserInfoParams struct {
	RequestId  string `json:"requestId"`
	SubmitTime string `json:"submitTime"`
	OrgId      string `json:"orgId"`
	Uid        string `json:"uid"`
}

type HmMemberUserInfoResult struct {
	OrgId    string `json:"orgId"`
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}