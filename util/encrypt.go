package util

import (
	"crypto/md5"
	"encoding/hex"
)

/*
md5加密
	before:加密前的字符串
	after:加密后的字符串
*/
func MD5(before string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(before))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
