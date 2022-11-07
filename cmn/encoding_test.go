package cmn

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_jwt(t *testing.T) {
	j := NewJWT("kkkkkkkkkdsadsakkk")
	kv := NewMapString().Put("user", "ssssssss").Put("group", "xxxxxxxxxx")
	tk, err := j.CreateToken(kv, -9) // 创建马上超时的令牌
	Info(tk, err)

	Info(j.Parse(tk))
	Info(j.Validate(tk))
	Info(j.IsExpired(tk))

	// 续签令牌
	tk, err = j.ExpandToken(tk, 5*time.Minute) // 续签5分钟超时的令牌
	Info(tk, err)
	Info(j.Parse(tk))
	Info(j.Validate(tk))
	Info(j.IsExpired(tk))
}

func Test_rsa(t *testing.T) {

	err := GenerateRsaKeyFile(2048, "f:/pri.pem", "f:/pub.pem")
	if err != nil {
		Error(err)
	}

	src := "公钥加密则私钥解密，私钥加密则公钥解密，本包只支持公钥加密私钥解密。公钥是公开的，私钥自己持有，通常私钥用于解密才具备秘钥的意义。"
	encodeStr, err := EncodeRsaByPubFile(src, "f:/pub.pem")
	if err != nil {
		Error(err)
	} else {
		Info(encodeStr)
	}

	str, err := DecodeRsaByPriFile(encodeStr, "f:/pri.pem")
	if err != nil {
		Error(err)
	} else {
		Info(str)
	}

	// 非文件方式
	byPri, byPub, err := GenerateRSAKey(2048)
	if err != nil {
		Error(err)
	}
	encodeStr, err = EncodeRsa(src, BytesToString(byPub))
	if err != nil {
		Error(err)
	} else {
		Info(encodeStr)
	}

	str, err = DecodeRsa(encodeStr, BytesToString(byPri))
	if err != nil {
		Error(err)
	} else {
		Info(str)
	}
}

func Test_base64(t *testing.T) {
	data := StringToBytes("这是待加密的字符串abc这是待加密的字符串abc这是待这是待加密的字符串abcxxx")
	s := Base64(data)

	Info(s)
	by, _ := Base64Decode(s)
	Info(BytesToString(by))
}

func Test_aes_ecb(t *testing.T) {
	src := "这是待加密的字符串abc这是待加密的字符串abc这是待这是待加密的字符串abc"
	key := "这是秘钥"
	aes := NewAesEcb(key)

	encode, _ := aes.Encode(src)
	log.Println((encode))

	decode, err := aes.Decode(encode)
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
