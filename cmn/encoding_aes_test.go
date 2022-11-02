package cmn

import (
	"log"
	"testing"
)

func Test_aes_ecb(t *testing.T) {
	src := "这是待加密的字符串abc这是待加密的字符串abc这是待这是待加密的字符串abc"
	key := "这是秘钥"
	aes := NewEncryptAes()

	encode, _ := aes.EncodeStr(src, key)
	log.Println((encode))

	decode, err := aes.DecodeStr(encode, key)
	log.Println((decode), err)

}
