package db

import (
	"github.com/mszsgo/himkt/env"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

const (
	COLLECTION_HM_VOUCHER_PROD = "hm_voucher_prod"
)

//券产品类型(1=满减券  2=满折券  3=礼物券 9=礼包券)
type EnumVoucherProdType string

const (
	VOUCHER_PROD_TYPE_1 EnumVoucherProdType = "1" // 满减券
	VOUCHER_PROD_TYPE_2 EnumVoucherProdType = "2" // 满折券
	VOUCHER_PROD_TYPE_3 EnumVoucherProdType = "3" // 兑换券
	VOUCHER_PROD_TYPE_9 EnumVoucherProdType = "9" // 礼包券
)

// 状态(0=新建 1=正常  2=停售  8=过期 9=删除)
type EnumVoucherProdStatus string

const (
	VOUCHER_PROD_STATUS_0 EnumVoucherProdStatus = "0"
	VOUCHER_PROD_STATUS_1 EnumVoucherProdStatus = "1"
	VOUCHER_PROD_STATUS_2 EnumVoucherProdStatus = "2"
	VOUCHER_PROD_STATUS_8 EnumVoucherProdStatus = "8"
	VOUCHER_PROD_STATUS_9 EnumVoucherProdStatus = "9"
)

// 是否外部导入券（1=是 0=否）
type EnumVoucherProdImport string

const (
	VOUCHER_PROD_IMPORT_Y EnumVoucherProdImport = "1"
	VOUCHER_PROD_IMPORT_N EnumVoucherProdImport = "0"
)

// 券产品账户
type HmVoucherProd struct {
	OrgId        string                `bson:"orgId"`        // 券产品所属机构编号
	OrgName      string                `bson:"orgName"`      // 券产品所属机构名称
	ProdId       string                `bson:"ProdId"`       // 产品编号
	Title        string                `bson:"title"`        // 产品标题名称
	Explain      string                `bson:"explain"`      // 产品简单说明
	Instructions string                `bson:"instructions"` //产品使用说明
	Type         EnumVoucherProdType   `bson:"type"`         // 券产品类型(1=满减券  2=满折券  3=兑换券 9=礼包券)
	Full         int64                 `bson:"full"`         // 满多少才能享受满减或者满折的金额，单位分
	FaceValue    int64                 `bson:"faceValue"`    // 满减面值，消费时可抵扣的金额，单位分 ,满折时代表折扣率，如 80 代表 8 折，计算时，金额*80%
	Pack         []string              `bson:"pack"`         // 礼包券包含的券产品编号
	Stock        int64                 `bson:"stock"`        // 剩余库存
	TotalSales   int64                 `bson:"totalSales"`   // 销售数量，发给用户的累计数量
	TotalUsage   int64                 `bson:"totalUsage" `  // 累计使用数量，核销数量
	Status       EnumVoucherProdStatus `bson:"status"`       // 状态(0=新建 1=正常  2=停售  8=过期 9=删除)
	Import       EnumVoucherProdImport `bson:"import"`       // 是否外部导入券（1=是 0=否）
	BegTime      time.Time             `bson:"begTime"`      // 有效期开始时间
	EndTime      time.Time             `bson:"endTime"`      // 有效期结束时间
	CreatedAt    time.Time             `bson:"createdAt"`    // 创建时间
	UpdatedAt    time.Time             `bson:"updatedAt"`    // 更新时间
	Remark       string                `bson:"remark"`       // 备注信息
}

func TestFind(t *testing.T) {
	c := Connect(env.HM_MONGO_CONNECTION_STRING).Collection("hm_voucher_prod")

	var prods []*HmVoucherProd
	err := Find(c, &prods, nil, bson.M{})
	if err != nil {
		t.Error(err)
	}
	t.Log(len(prods))
	t.Log(prods)
}
