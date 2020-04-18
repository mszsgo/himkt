package env

import "os"

// 读取系统环境变量
var (
	HM_ENV                     = os.Getenv("HM_ENV") //可选值：local/dev/uat/prd
	HM_HTTP_PROXY              = os.Getenv("HM_HTTP_PROXY")
	HM_MONGO_CONNECTION_STRING = os.Getenv("HM_MONGO_CONNECTION_STRING")
)
