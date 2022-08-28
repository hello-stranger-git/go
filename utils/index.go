package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	return MD5Bytes([]byte(s))
}
func MD5Bytes(s []byte) string {
	ret := md5.Sum(s)
	return hex.EncodeToString(ret[:])
}
