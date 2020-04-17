package cfg

// 阿里云对象存储

type ossCfg struct {
}

func (c *Cfg) Oss() *ossCfg {
	return &ossCfg{}
}
