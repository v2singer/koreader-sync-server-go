package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hex(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
} 