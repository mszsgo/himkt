package src

import (
	"context"
	"github.com/mszsgo/himkt/micro"
	"github.com/mszsgo/himkt/time/t14"
)

func HmUnionpayQualReduce(ctx context.Context, params *HmUnionpayQualReduceParams) (r *HmUnionpayQualReduceResult, err error) {
	params.SubmitTime = t14.NowF14()
	err = micro.Call(ctx, "hm.unionpay.qual.reduce", params, &r)
	return
}

type HmUnionpayQualReduceParams struct {
	RequestId      string `json:"requestId"`
	SubmitTime     string `json:"submitTime"`
	MchId          string `json:"mchId" description:"接入商户编号"`
	TransNumber    string `json:"transNumber" description:"流水号，确保唯一，业务系统可使用订单号、券编号等"`
	QualNum        string `json:"qualNum" description:"资格池编号,云闪付申请表上有填写"`
	QualType       string `json:"qualType" description:"资格类型   固定值“open_id”、“mobile”、“card_no”"`
	QualValue      string `json:"qualValue" description:"资格类型对应的值，如手机号"`
	ActivityNumber string `json:"activityNumber" description:"云闪付提供的活动编号"`
}

type HmUnionpayQualReduceResult struct {
	RespCode    string                               `json:"respCode"`
	RespTime    string                               `json:"respTime"`
	TransNumber string                               `json:"transNumber"`
	AwardInfo   *HmUnionpayQualReduceResultAwardInfo `json:"awardInfo"`
}

type HmUnionpayQualReduceResultAwardInfo struct {
	ActivityNumber  string `json:"activityNumber"`
	ActivityName    string `json:"activityName"`
	BeginTime       string `json:"beginTime"`
	EndTime         string `json:"endTime"`
	AwardId         string `json:"awardId"`
	AwardType       string `json:"awardType"`
	AwardName       string `json:"awardName"`
	ExtAcctId       string `json:"extAcctId"`
	ExtAcctName     string `json:"extAcctName"`
	DrawDesc        string `json:"drawDesc"`
	CouponStartDate string `json:"couponStartDate"`
	CouponEndDate   string `json:"couponEndDate"`
	CouponGoodsUrl  string `json:"couponGoodsUrl"`
}
