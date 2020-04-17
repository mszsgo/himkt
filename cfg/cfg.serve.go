package cfg

type serveCfg struct {
	Host string
	Port int64
}

func (c *Cfg) Serve() *serveCfg {
	return &serveCfg{
		Port: 80,
	}
}
