package hm

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mszsgo/himkt/protocol"
	"io/ioutil"
	"net/http"
	"strings"
)

type Error500 struct {
	Errno string `json:"errno"`
	Error string `json:"error"`
}

func Do(ctx context.Context, method string, body string) ([]byte, error) {
	request, err := http.NewRequest("POST", "http://"+strings.Split(method, ".")[1]+"/"+method, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var client = protocol.HttpClient()

	// ctx 不等于空时，读取header与client
	if ctx != nil {
		ctxTrack := ctx.Value("track")
		if ctxTrack != nil {
			track := ctxTrack.(*Track)
			request.Header.Set("sid", track.Sid)
			request.Header.Set("pid", track.Tid)
			request.Header.Set("tid", uuid.New().String())
		}

		ctxClient := ctx.Value("client")
		if ctxClient != nil {
			client = ctxClient.(*http.Client)
		}
	}
	response, err := client.Do(request)
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
	return nil, errors.New("99999:Gateway->API异常(" + response.Status + ")，请检查业务服务")
}

func Call(ctx context.Context, method string, i interface{}, o interface{}) (err error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		return err
	}
	resp, err := Do(ctx, method, string(bytes))
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, o)
	return err
}
