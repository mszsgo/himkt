package token

import (
	"github.com/mszsgo/himkt/errs"
	"github.com/mszsgo/himkt/genid"
	"github.com/mszsgo/himkt/rdb"
)

/*
 toa/tob/toc   接口token设置与校验

*/

// Redis KeyFix
const (
	RKEY_TOKEN rdb.KeyPre = "token:" // 用户会话Token ,key为uuid   value为用户id
)

//生成token
func Gen(uid string, validSecond int64) (token string, err error) {
	token = genid.UUID()
	err = RKEY_TOKEN.Set(token, uid, validSecond)
	return
}

// 根据token获取用户id
func Get(token string) (uid string) {
	return RKEY_TOKEN.Get(token)
}

// 验证token,判断token对应的内容是否等于用户编号
func Verify(token, uid string) error {
	v := RKEY_TOKEN.Get(token)
	if v != "" && v != uid {
		return errs.E99911
	}
	return nil
}
