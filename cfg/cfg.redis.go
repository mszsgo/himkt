package cfg

type redisCfg struct {
	Url string `json:"url"`
}

func Redis() *redisCfg {
	var redisCfg *redisCfg
	NowConfig("redis", &redisCfg)
	return redisCfg
}
