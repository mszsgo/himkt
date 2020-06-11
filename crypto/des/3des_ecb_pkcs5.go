package des

import (
	"crypto/des"
	"encoding/base64"
)

// 加密结果以base64输出

//ECB加密
func TRIPLE_DES_ECB_PKCS5_Encode(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewTripleDESCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out)
}

//ECB解密
func TRIPLE_DES_ECB_PKCS5_Decode(src, key string) string {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		panic(err)
	}
	keyByte := []byte(key)
	block, err := des.NewTripleDESCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}
