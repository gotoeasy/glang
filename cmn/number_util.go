package cmn

import (
	"encoding/binary"
	"math/rand"
	"strconv"
	"time"
)

// int 转 string
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

// uint32 转 []byte
func Uint32ToBytes(num uint32) []byte {
	bkey := make([]byte, 4)
	binary.BigEndian.PutUint32(bkey, num)
	return bkey
}

// []byte 转 uint32
func BytesToUint32(bytes []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(bytes))
}

// 随机数
func RandomInt(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
