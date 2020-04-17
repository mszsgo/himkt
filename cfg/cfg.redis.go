package cfg

type redisCfg struct {
}

func (c *Cfg) Redis() *redisCfg {
	return &redisCfg{}
}
