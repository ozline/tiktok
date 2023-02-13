package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 use md5 to encode string
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
