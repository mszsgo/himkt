package cfg

type HimktCfg struct {
	Domain string `json:"domain" description:"平台服务基础访问域名，仅包含协议与域名，如：https://ms.himkt.cn"`
	Server struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"server"`
}

func Himkt() *HimktCfg {
	var hm *HimktCfg
	NowConfig("himkt", &hm)
	if hm.Server.Port == 0 {
		hm.Server.Port = 80
	}
	return hm
}
