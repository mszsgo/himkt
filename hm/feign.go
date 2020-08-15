package hm

import (
	"context"
	"encoding/json"
	"errors"
	"himkt/genid"
	"himkt/protocol"
	"io/ioutil"
	"net/http"
	"strings"
)

type Error500 struct {
	Errno string `json:"errno"`
	Error string `json:"error"`
}

func Post(ctx context.Context, url string, body string) ([]byte, error) {
	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var client = protocol.HttpClient()

	// ctx 不等于空时，读取header与client
	if ctx != nil {
		request.Header.Set("appid", ctx.Value("appid").(string))

		ctxTrack := ctx.Value("track")
		var track *Track
		if ctxTrack != nil {
			track = ctxTrack.(*Track)
		} else {
			track = &Track{
				Sid: genid.UUID(),
				Pid: genid.UUID(),
			}
		}
		request.Header.Set("sid", track.Sid)
		request.Header.Set("pid", track.Tid)

		ctxClient := ctx.Value("client")
		if ctxClient != nil {
			client = ctxClient.(*http.Client)
		}
	}
	request.Header.Set("tid", genid.UUID())
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
	return nil, errors.New("99999:HTTP异常(" + response.Status + ")，请检查业务服务 POST " + url)
}

func Do(ctx context.Context, method string, body string) ([]byte, error) {
	name := strings.Split(method, ".")[1]
	url := "http://" + name + "/" + method
	return Post(ctx, url, body)
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
