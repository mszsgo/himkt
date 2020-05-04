package cfg

import (
	"encoding/json"
	"fmt"
	"github.com/mszsgo/himkt/env"
	"github.com/mszsgo/himkt/protocol"
	"io/ioutil"
)

// 加载配置
func LoadConfig(name string) ([]byte, error) {
	// 如果是local本地开发环境，配置服务名加后缀`-local`，仅用于本地开发环境，发布无需配置此环境变量
	if env.HM_ENV == "local" {
		name = name + "-local"
	}
	resp, err := protocol.HttpClient().Get("http://config/get?name=" + name)
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
		fmt.Printf("Error:加载配置"+name+"失败 %s", err.Error())
		panic(err)
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		fmt.Printf("Error:加载配置"+name+"失败 %s", err.Error())
		panic(err)
	}
	return
}
