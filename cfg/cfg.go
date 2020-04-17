package cfg

import (
	"encoding/json"
	"github.com/mszsgo/himkt"
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
				if himkt.HM_HTTP_PROXY != "" {
					request.URL.Host = himkt.HM_HTTP_PROXY
				}
				return request.URL, err
			},
		},
		Timeout: 5 * time.Second,
	}
	// 如果是local本地开发环境，配置服务名加后缀`-local`，仅用于本地开发环境，发布无需配置此环境变量
	if himkt.HM_ENV == "local" {
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
		panic(err)
		return
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		panic(err)
		return
	}
	return
}
