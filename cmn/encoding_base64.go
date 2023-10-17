package cmn

import (
	"encoding/base64"
)

// Base64编码（同Base64Encode）
func Base64(bts []byte) string {
	return base64.StdEncoding.EncodeToString(bts)
}

// Base64编码
func Base64Encode(bts []byte) string {
	return base64.StdEncoding.EncodeToString(bts)
}

// Base64解码
func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
