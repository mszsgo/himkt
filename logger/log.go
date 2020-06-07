package logger

import (
	"github.com/sirupsen/logrus"
)

// 默认日志，使用阿里云SLS日志服务
func Now(topic string) (l *logrus.Logger) {
	defer func() {
		l = logrus.StandardLogger()
	}()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Add 阿里云日志服务Hook ,注意日志输出格式只能使用JSON
	logCfg := LogCfg()
	if logCfg == nil {
		logrus.Warn("未读取到log配置信息")
		return
	}
	switch logCfg.Level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	if logCfg.Aliyun != nil {
		sls := logCfg.Aliyun
		sls.topic = topic
		logrus.AddHook(&AliyunSlsHook{AliyunSls: sls})
		logrus.Info("add aliyun sls hook ")
	}
	return
}
