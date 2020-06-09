package cfg

type serveCfg struct {
	Host string
	Port int64
}

// @deprecated
func Serve() *serveCfg {
	return &serveCfg{
		Port: 80,
	}
}
