package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

func Encrypt(data, key []byte, mode string, padding string) {

}

func Decrypt() {

}

type Padding string

const (
	DES_CBC_PKCS7 Padding = "DES/ECB/PKCS7Padding"
)

//明文补码算法
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padtext...)
}

//明文减码算法
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// 加密,结果转为Base64字符串
func EncryptCBC(origData, key []byte) (string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	origData = PKCS7Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// 解密,结果转为明文字符串
func DecryptCBC(origData string, key []byte) (string, error) {
	origBytes, err := base64.StdEncoding.DecodeString(origData)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	crypted := make([]byte, len(origBytes))
	blockMode.CryptBlocks(crypted, origBytes)
	crypted = PKCS7UnPadding(crypted)
	return string(crypted), nil
}
