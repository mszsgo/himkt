package protocol

import (
	"himkt/env"
	"net/http"
	"net/url"
	"time"
)

func HttpClient() *http.Client {
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
}
