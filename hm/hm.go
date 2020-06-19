package hm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	ERR_HTTP_METHOD = errors.New("99999:请使用HTTP POST请求")
)

func ResponseSuccess(i interface{}) []byte {
	if i == nil {
		return []byte("")
	}
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	suc := `{"errno":"00000","error":"ok"}`
	if len(bytes) > 2 {
		bytes = append([]byte(suc[0:len(suc)-1]+","), bytes[1:]...)
	} else {
		bytes = []byte(suc)
	}
	return bytes
}

func ResponseFail(err error) []byte {
	msg := err.Error()
	if msg[5:6] == ":" {
		return []byte(fmt.Sprintf(`{"errno":"%s","error":"%s"}`, msg[0:5], msg[6:]))
	} else {
		return []byte(fmt.Sprintf(`{"errno":"%s","error":"%s"}`, "99999", msg))
	}
}

func ResponseWriter(writer http.ResponseWriter, status int, content []byte) {
	writer.Write(content)
	if status != 200 {
		writer.WriteHeader(status)
	}
}

func post(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		if request.Method != "POST" {
			ResponseWriter(writer, 500, ResponseFail(ERR_HTTP_METHOD))
			return
		}
		handle.ServeHTTP(writer, request)
	})
}

// 服务调用轨迹记录
type Track struct {
	Sid string `json:"sid" description:"会话编号"`
	Pid string `json:"pid" description:"上级编号"`
	Tid string `json:"tid" description:"交易编号"`
}

type ResolveParams struct {
	Context context.Context
	Request *http.Request
	Writer  http.ResponseWriter
	Body    []byte
}

func (p *ResolveParams) BodyUnmarshal(i interface{}) {
	err := json.Unmarshal(p.Body, i)
	if err != nil {
		panic(err)
	}
}

// 定义接口，注意：pattern必须/开头
func DefApi(pattern string, resolve func(p *ResolveParams) (out interface{}, err error)) {
	http.Handle(pattern, post(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		begTime := time.Now().UnixNano() / 1e6

		h := request.Header
		track := &Track{Sid: h.Get("sid"), Pid: h.Get("pid"), Tid: h.Get("tid")}
		hlog := log.WithField("sid", track.Sid).WithField("pid", track.Pid).WithField("tid", track.Tid).WithField("method", pattern[1:])

		defer func() {
			if e := recover(); e != nil {
				var err error
				switch v := e.(type) {
				case error:
					err = e.(error)
				case *log.Entry:
					err = errors.New(e.(*log.Entry).Message)
				case string:
					err = errors.New(e.(string))
				default:
					err = errors.New(fmt.Sprintf("ERROR:未知错误类型 %v ", v))
				}
				resBytes := ResponseFail(err)
				hlog.WithField("responseBody", string(resBytes)).Info("响应报文")
				ResponseWriter(writer, 500, resBytes)
			}
			endTime := time.Now().UnixNano() / 1e6
			hlog.WithField("milliseconds", strconv.FormatInt(endTime-begTime, 10)).Info("毫秒")
		}()

		// 调用业务方法
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		hlog.WithField("requestBody", string(body)).Info("请求报文")
		i, e := resolve(&ResolveParams{Body: body, Request: request, Writer: writer, Context: context.WithValue(context.Background(), "track", track)})
		if e != nil {
			panic(e)
		}
		resBytes := ResponseSuccess(i)
		hlog.WithField("responseBody", string(resBytes)).Info("响应报文")
		ResponseWriter(writer, 200, resBytes)
	})))
}

// 接口请求Model
type RequestHeader struct {
	Appid      string `json:"appid"`
	Method     string `json:"method"`
	RequestId  string `json:"requestId"`
	SubmitTime string `json:"submitTime"`
}

// 接口响应 Model
type ResponseHeader struct {
	Errno      string `json:"errno"`
	Error      string `json:"error"`
	HostTime   string `json:"hostTime"`
	HostNo     string `json:"hostNo"`
	Appid      string `json:"appid"`
	Method     string `json:"method"`
	RequestId  string `json:"requestId"`
	SubmitTime string `json:"submitTime"`
}
