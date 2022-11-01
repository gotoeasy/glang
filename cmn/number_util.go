package cmn

import (
	"encoding/binary"
	"strconv"
)

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Uint32ToBytes(num uint32) []byte {
	bkey := make([]byte, 4)
	binary.BigEndian.PutUint32(bkey, num)
	return bkey
}

func BytesToUint32(bytes []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(bytes))
}
