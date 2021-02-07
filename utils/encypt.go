package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * md5加密字符串
 * @param string str 明文字符串
 * @return string
 */
func Md5String(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}
