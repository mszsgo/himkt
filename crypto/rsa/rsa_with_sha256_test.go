package rsa

import (
	"crypto/rsa"
	"encoding/base64"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"testing"
)

const priKey = `-----BEGIN RSA PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAN0lpVP6NzuqOVfH
MrLanvozMCbYBoSZvDSOIJ1RrzLczYbFJAg8Uk0Ez1iQOPYkrWQMsrm60CpETN+y
22LVbT4yAmk+xQe9JxrxIKZOIte172EFoo94x4+0kxvirgY533BfEfh35kFYVMbN
zqKgLLKUKI16QvPhRd5bnEVJut3NAgMBAAECgYBDggh0GD/gINicaRya6I37RsD+
OpfSxrNjP7fJDnRzDGBZMbkMIo3lIQDCwlQBH7Umg5HTjKmbHXpyF0FNbGWKM1d8
T2T9I9d+VNflayyNd+B232laV6Fd3Yd4AgKEidnNZopYdc5Lv22NDBnVIe84Q5zo
ggkPLmtVj/M9W3I9KQJBAP4/f18QNINK3Aw6ep5T720PMuhkjnwlxkAgdsRUj13K
lKLr6scTHjoDcedlmwqS689pV07uwi9PcP9NJH5SuLMCQQDeq8HI9rfvGNiY6Z9M
llCERr6cLV87FLp/W9KTwqRIBolFqNFD63SiLaZdZz97Ijtu672eCgmfz6EknCTu
hk9/AkAsb3CVbsGeyLCYuoe/sC/CQcvF7f0xGXECuCnJunWsEMHVj6Hi1SaNQRjh
NlLYkEECaQjHVNc/aYT4OcVNBqwdAkEArO+ygnTzJ5QhbKlNGjtH72POWBd49bbE
nXNgGXg7BOFmrZ5OYLS2kdpocfRlPim5BQ2758kd7mtgWMGYgVP42wJBALghOyl4
Psy9Xcvrlk4I+AWlLXRb4pmpgapBTO0o+cdpoEOxV9idVcSWfCrPCDEewmw/49QW
gBKFLya3nOO6Muo=
-----END RSA PRIVATE KEY-----
`

const pubKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDdJaVT+jc7qjlXxzKy2p76MzAm
2AaEmbw0jiCdUa8y3M2GxSQIPFJNBM9YkDj2JK1kDLK5utAqREzfstti1W0+MgJp
PsUHvSca8SCmTiLXte9hBaKPeMePtJMb4q4GOd9wXxH4d+ZBWFTGzc6ioCyylCiN
ekLz4UXeW5xFSbrdzQIDAQAB
-----END PUBLIC KEY-----
`

func TestSign(t *testing.T) {
	value := "123"

	sign, err := RsaWithSha256Sign(value, priKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("sign=" + sign)

	err = RsaWithSha256Verify(value, sign, pubKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("verify success")
}

// 解析 pfx 证书私钥与序列号
func TestPfxToBase64(t *testing.T) {
	fbytes, _ := ioutil.ReadFile("D:\\Information\\Company\\中国银联\\手机网页支付\\手机网页（WAP）支付产品技术开发包1.1.8\\Java Version SDK (通用版)\\ACPSample_B2C\\src\\assets\\测试环境证书\\acp_test_sign.pfx")
	pass := "000000"

	pfxBase64 := base64.StdEncoding.EncodeToString(fbytes)
	t.Log(pfxBase64)

	priKey, cert, err := pkcs12.Decode(fbytes, pass)
	if err != nil {

	}
	t.Log(priKey.(*rsa.PrivateKey))
	t.Log("SerialNumber=" + cert.SerialNumber.String())
}
