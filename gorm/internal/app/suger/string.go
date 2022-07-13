package suger

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
)

// 字符串拼接
func StringContact(item ...string) string {
	var itemBuffer bytes.Buffer
	for _, v := range item {
		itemBuffer.WriteString(v)
	}
	return itemBuffer.String()
}

//获取字符串的md5
func GetMd5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
