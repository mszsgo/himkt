package rsa

import (
	"crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 签名
func RsaWithSha256Sign(value string, priKey string) (sign string, err error) {
	if priKey == "" {
		return "", errors.New("ERROR:RSA Sign 商户私钥不能为空")
	}

	p, _ := pem.Decode([]byte(priKey))
	privateKey, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil {
		return
	}

	hash := sha256.New()
	hash.Write([]byte(value))
	shaBytes := hash.Sum(nil)
	b, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, shaBytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
}

// 验签
func RsaWithSha256Verify(value string, sig string, pubKey string) (err error) {
	if pubKey == "" {
		return errors.New("ERROR:RSA Sign 银联公钥不能为空")
	}
	oriSign, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return
	}

	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		err = errors.New("public key error")
		return
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	hashed := sha256.Sum256([]byte(value))
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], oriSign)
	if err != nil {
		return
	}
	return
}
