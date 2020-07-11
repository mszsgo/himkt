package t14

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	HHMMSS = "150405"
)

func NowF6() string {
	return Time6(time.Now()).String()
}

type Time6 time.Time

func (nt Time6) String() string {
	return time.Time(nt).Local().Format(HHMMSS)
}

func (nt Time6) MarshalJSON() ([]byte, error) {
	t := time.Time(nt).Local()

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(HHMMSS)+2)
	b = append(b, '"')
	if !t.IsZero() {
		b = t.AppendFormat(b, HHMMSS)
	}
	b = append(b, '"')
	return b, nil
}

func (nt *Time6) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), "\"")
	if data == nil || v == "" {
		*nt = Time6(time.Time{})
		return nil
	}
	t, err := time.Parse(HHMMSS, v)
	if err != nil {
		return fmt.Errorf("解析时间字符串'%s'出错", v)
	}
	*nt = Time6(t.Local())
	return err
}
