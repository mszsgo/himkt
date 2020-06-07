package md5

import (
	"crypto/md5"
	"fmt"
)

func Encrypt(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}
