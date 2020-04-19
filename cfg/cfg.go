package cfg

import (
	"encoding/json"
	"github.com/mszsgo/himkt/env"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// 加载配置
func LoadConfig(name string) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			// 服务HTTP代理配置，示例系统环境变量： "MS_HTTP_PROXY=211.152.57.29:39084"
			Proxy: func(request *http.Request) (url *url.URL, err error) {
				if env.HM_HTTP_PROXY != "" {
					request.URL.Host = env.HM_HTTP_PROXY
				}
				return request.URL, err
			},
		},
		Timeout: 5 * time.Second,
	}
	// 如果是local本地开发环境，配置服务名加后缀`-local`，仅用于本地开发环境，发布无需配置此环境变量
	if env.HM_ENV == "local" {
		name = name + "-local"
	}
	resp, err := client.Get("http://config/get?name=" + name)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func NowConfig(name string, v interface{}) {
	bytes, err := LoadConfig(name)
	if err != nil {
		log.Error("加载配置"+name+"失败", err.Error())
		panic(err)
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		log.Error("加载配置"+name+"失败", err.Error())
		panic(err)
	}
	return
}
