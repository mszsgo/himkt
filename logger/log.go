package logger

import (
	"github.com/sirupsen/logrus"
)

// 默认日志，使用阿里云SLS日志服务
func New(topic string) *logrus.Logger {
	var logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	// Add 阿里云日志服务Hook ,注意日志输出格式只能使用JSON
	sls := LogCfg().Aliyun
	sls.topic = topic
	logger.AddHook(&aliyunSlsHook{AliyunSls: sls})

	return logger
}
