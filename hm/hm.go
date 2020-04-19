package hm

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

var (
	ERR_HTTP_METHOD = errors.New("99999:请使用HTTP POST请求")
)

func ResponseSuccess(writer http.ResponseWriter, i interface{}) {
	writer.WriteHeader(200)
	if i == nil {
		return
	}
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	writer.Write(bytes)
}

func ResponseFail(writer http.ResponseWriter, err error) {
	writer.WriteHeader(500)
	msg := err.Error()
	if msg[5:6] == ":" {
		writer.Write([]byte(fmt.Sprintf(`{"errno":"%s","error":"%s"}`, msg[0:5], msg[6:])))
	} else {
		writer.Write([]byte(fmt.Sprintf(`{"errno":"%s","error":"%s"}`, "99999", msg)))
	}
}

func post(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		if request.Method != "POST" {
			ResponseFail(writer, ERR_HTTP_METHOD)
			return
		}
		handle.ServeHTTP(writer, request)
	})
}

type ResolveParams struct {
	request *http.Request
	writer  http.ResponseWriter
}

func (p *ResolveParams) BodyUnmarshal(i interface{}) {
	data, err := ioutil.ReadAll(p.request.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, i)
	if err != nil {
		panic(err)
	}
}

func DefApi(pattern string, resolve func(p *ResolveParams) (out interface{}, err error)) {
	// TODO pattern值非/开头的，拼接/符号

	http.Handle(pattern, post(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
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
					err = errors.New(fmt.Sprintf("99999：未知错误类型 %v ", v))
				}
				ResponseFail(writer, err)
			}
		}()
		i, e := resolve(&ResolveParams{request: request, writer: writer})
		if e != nil {
			panic(e)
		}
		ResponseSuccess(writer, i)
	})))
}
