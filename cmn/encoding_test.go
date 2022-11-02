package cmn

import (
	"fmt"
	"log"
	"testing"
)

func Test_aes_ecb(t *testing.T) {
	src := "这是待加密的字符串abc这是待加密的字符串abc这是待这是待加密的字符串abc"
	key := "这是秘钥"
	aes := NewAesEcb()

	encode, _ := aes.EncodeStr(src, key)
	log.Println((encode))

	decode, err := aes.DecodeStr(encode, key)
	log.Println((decode), err)

}

func Test_tostring(t *testing.T) {
	s := "f"
	// f("0aaaaaaaaa啊aaads")
	for i := 1; i < 100; i++ {
		s += fmt.Sprintf("%v", i)
		log.Println(HashDJB(s))
		//log.Println(HashCode(StringToBytes(s)))
	}

}
func DJB(str string) uint32 {
	r := []rune(str)
	var hash uint32 = 5381
	for i := 0; i < len(r); i++ {
		hash += (hash << 5) ^ uint32(r[i])
	}
	return hash
}

func HashDJB(str string) uint32 {
	var rs uint32 = 53653 // 5381
	r := []rune(str)
	for i := len(r) - 1; i >= 0; i-- {
		rs = (rs * 33) ^ uint32(r[i])
	}
	return rs
}
