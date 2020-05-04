package logger

import "github.com/mszsgo/himkt/cfg"

// 阿里云日志服务接入

type logCfg struct {
	Aliyun *AliyunSls `json:"aliyun"`
}

func LogCfg() *logCfg {
	var l *logCfg
	cfg.NowConfig("log", &l)
	return l
}
