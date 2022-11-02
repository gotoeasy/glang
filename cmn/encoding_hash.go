package cmn

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

func HashCode(bts []byte) uint32 {
	return crc32.ChecksumIEEE(bts)
}

func Hash(str string) uint32 {
	var rs uint32 = 53653 // 5381
	r := []rune(str)
	for i := len(r) - 1; i >= 0; i-- {
		rs = (rs * 33) ^ uint32(r[i])
	}
	return rs
}

// 字符串哈希处理后取模(余数)，返回值最大不超过mod值
func HashMod(str string, mod uint32) string {
	return fmt.Sprint(Hash("添油"+str+"加醋") % mod)
}

func Base64(bts []byte) string {
	return base64.StdEncoding.EncodeToString(bts)
}

func Base64Encode(bts []byte) string {
	return base64.StdEncoding.EncodeToString(bts)
}

func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func Md5(bts []byte) string {
	h := md5.New()
	h.Write(bts)
	return hex.EncodeToString(h.Sum(nil))
}
