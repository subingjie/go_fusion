package upload

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

type TencentCos struct {
}

/*
//  keyStr 密钥
//  value  消息内容
*/
func HMACSHA1(keyStr, value string) string {

	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(value))
	//进行base64编码
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return res
}
