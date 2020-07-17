package hmnotify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mszsgo/himkt/errs"
	"github.com/mszsgo/himkt/hm"
	"github.com/mszsgo/himkt/micro/hmopen"
	log "github.com/sirupsen/logrus"
	"net/url"
)

// 通知接口
// 此处只处理通知请求操作，重试逻辑由业务系统使用定时任务控制
// 返回nil ，代表处理成功，否则通知失败需要重试
func Request(ctx context.Context, serviceUrl, appid, method string, bizObj interface{}) (err error) {
	plog := log.WithField("appid", appid).WithField("method", method).WithField("serviceUrl", serviceUrl)
	defer func() {
		if err != nil {
			plog.Warn("通知请求异常 " + err.Error())
		}
	}()

	plog.Info("调用通知接口")
	bizBytes, err := json.Marshal(bizObj)
	if err != nil {
		return
	}
	bizJson := string(bizBytes)
	plog.Info("通知请求报文：" + bizJson)

	// TODO 从数据表读取秘钥，待优化为缓存方案，在更新秘钥时同时更新缓存
	app, err := hmopen.GetOpenApp(appid, false)
	if err != nil {
		return
	}

	// 加密
	reqBiz, err := hm.EncryptDesMd5(bizJson, app.DesKey)
	if err != nil {
		return
	}
	plog.Debug("通知请求密文：" + reqBiz)

	// HTTP 请求
	body := fmt.Sprintf("method=%s&appid=%s&biz=%s", method, appid, url.QueryEscape(reqBiz))
	resBytes, err := hm.Post(ctx, serviceUrl, body)
	if err != nil {
		return
	}
	resBody := string(resBytes)
	plog.Debug("通知响应密文：" + resBody)

	//解密
	resBiz, err := hm.DecryptDesMd5(resBody, app.DesKey)
	if err != nil {
		err = errors.New("99999:响应报文解密出错")
		return
	}
	plog.Info("通知响应报文：" + resBiz)

	var mbiz map[string]interface{}
	err = json.Unmarshal([]byte(resBiz), &mbiz)
	if err != nil {
		return
	}
	if mbiz["errno"].(string) != errs.SUCCESS.Code {
		err = errors.New(mbiz["errno"].(string) + ":" + mbiz["error"].(string))
		return
	}
	return nil
}
