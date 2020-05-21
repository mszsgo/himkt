package logger

import (
	"github.com/sirupsen/logrus"
)

// 默认日志，使用阿里云SLS日志服务
func New(topic string) (logger *logrus.Logger) {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	// Add 阿里云日志服务Hook ,注意日志输出格式只能使用JSON
	logCfg := LogCfg()
	if logCfg == nil {
		logger.Warn("未读取到log配置信息")
		return
	}
	if logCfg.Aliyun != nil {
		sls := logCfg.Aliyun
		sls.topic = topic
		logger.AddHook(&AliyunSlsHook{AliyunSls: sls})
		logger.Info("add aliyun sls hook ")
	}
	return
}
