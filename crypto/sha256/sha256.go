package sha256

import (
	"crypto/sha256"
	"fmt"
)

// Sha256 转16进制输出字符串
func EncodeToHex(v string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(v)))
}
