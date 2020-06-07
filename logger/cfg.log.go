package logger

import (
	"fmt"
	"github.com/mszsgo/himkt/cfg"
)

// 阿里云日志服务接入

type logCfg struct {
	Level  string     `json:"level"` // debug  info  warn  error
	Aliyun *AliyunSls `json:"aliyun"`
}

func LogCfg() *logCfg {
	var l *logCfg
	err := cfg.NowConfig("log", &l)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return nil
	}
	if l.Level == "" {
		l.Level = "info"
	}
	return l
}
