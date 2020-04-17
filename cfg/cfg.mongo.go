package cfg

type mongoCfg struct {
	Url string `json:"url"`
}

func Mongo() *mongoCfg {
	var mongoCfg *mongoCfg
	NowConfig("mongo", &mongoCfg)
	return mongoCfg
}
