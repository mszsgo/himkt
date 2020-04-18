package hm

import (
	"encoding/json"
	"errors"
	"github.com/mszsgo/himkt/env"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Feign struct {
	client *http.Client
}

var CtxFeign = &Feign{
	client: func() *http.Client {
		if env.HM_HTTP_PROXY == "" {
			return http.DefaultClient
		}
		return &http.Client{
			Transport: &http.Transport{
				Proxy: func(request *http.Request) (url *url.URL, err error) {
					request.URL.Host = env.HM_HTTP_PROXY
					return request.URL, err
				},
			},
			Timeout: 5 * time.Second,
		}
	}(),
}

// 用于网关调用接口
func (df *Feign) Do(method string, body string, token string) ([]byte, error) {
	request, err := http.NewRequest("POST", "http://"+strings.Split(method, ".")[1]+"/"+method, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("AccessToken", token)
	response, err := df.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 200 {
		return ioutil.ReadAll(response.Body)
	}
	if response.StatusCode == 500 {
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		var e500 *Error500
		err = json.Unmarshal(bytes, &e500)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(e500.Errno + ":" + e500.Error)
	}
	return nil, errors.New("99999:Gateway->API异常HTTP状态码")
}

type Error500 struct {
	Errno string `json:"errno"`
	Error string `json:"error"`
}

// 调用接口并解析，用于微服务调用
func (df *Feign) Call(method string, i interface{}, o interface{}) (err error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		return err
	}
	resp, err := df.Do(method, string(bytes), "")
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, o)
	return err
}
