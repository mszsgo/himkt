package crypto

import (
	"crypto/sha256"
	"fmt"
)

// Sha256 z转16进制输出字符串
func Sha256(v string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(v)))
}
