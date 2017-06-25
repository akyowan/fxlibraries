package stringhelper

import (
	"bytes"
	"fmt"
	"github.com/satori/go.uuid"
)

//GererateHash 生成指定长度的随机hash字符串
func GererateHash(length int) string {
	var hash string
	for len(hash) < length {
		hash += GenerateUUIDString()
	}
	return hash[0:length]
}

//GenerateUUIDString 生成32位uuid字符串
func GenerateUUIDString() string {
	u := uuid.NewV4().Bytes()
	str := ByteToHex(u)
	return str
}

//ByteToHex 字符数组转十六进制字符串
func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {
		buffer.WriteString(fmt.Sprintf("%02x", int64(b&0xff)))
	}

	return buffer.String()
}
