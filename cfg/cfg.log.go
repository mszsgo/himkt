package cfg

// 阿里云日志服务接入

type logCfg struct {
}

func (c *Cfg) Log() *logCfg {
	return &logCfg{}
}
