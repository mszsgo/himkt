package errs

import (
	"encoding/json"
	"strings"
)

type Err struct {
	Code string `json:"errno"`
	Msg  string `json:"error"`
}

func New(code, msg string) *Err {
	return &Err{Code: code, Msg: msg}
}

func NewF(err error) *Err {
	e := err.Error()
	i := strings.Index(e, ":")
	if i == -1 {
		return &Err{Code: "ERROR", Msg: e}
	}
	return &Err{Code: e[0:i], Msg: e[i+1:]}
}

func (e *Err) Error() string {
	return e.Code + ":" + e.Msg
}

func (e *Err) JsonMarshal() []byte {
	b, _ := json.Marshal(e)
	return b
}