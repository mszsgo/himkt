package t14

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	DDHHMMSS = "150405"
)

type Time6 time.Time

func (nt Time6) String() string {
	return time.Time(nt).Local().Format(DDHHMMSS)
}

func (nt Time6) MarshalJSON() ([]byte, error) {
	t := time.Time(nt).Local()

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DDHHMMSS)+2)
	b = append(b, '"')
	if !t.IsZero() {
		b = t.AppendFormat(b, DDHHMMSS)
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
	t, err := time.Parse(DDHHMMSS, v)
	if err != nil {
		return fmt.Errorf("解析时间字符串'%s'出错", v)
	}
	*nt = Time6(t.Local())
	return err
}
