package cmn

import (
	"bytes"
	"crypto/aes"
	"errors"
)

type AesEcb struct{}

func NewAesEcb() *AesEcb {
	return &AesEcb{}
}

// 字符串加密
func (a *AesEcb) EncodeStr(src string, secret string) (string, error) {
	by, err := a.Encode(StringToBytes(src), StringToBytes(secret))
	if err != nil {
		return "", err
	}
	return Base64(by), nil
}

// 字符串解密
func (a *AesEcb) DecodeStr(src string, secret string) (string, error) {
	srcBy, err := Base64Decode(src)
	if err != nil {
		return "", err
	}

	by, err := a.Decode(srcBy, StringToBytes(secret))
	if err != nil {
		return "", err
	}

	return BytesToString(by), nil
}

func (a *AesEcb) Encode(src []byte, secret []byte) ([]byte, error) {
	if len(src) == 0 {
		return src, nil // 空内容加密结果仍旧空
	}

	appendkeys := StringToBytes("秘钥长度仅支持16位、24位、32位，如果参数的秘钥有误，则按默认补足至32位方便使用")
	if !checkKey(secret) {
		secret = append(secret, appendkeys...)[:32]
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}
	paddingSrc := aesPkcs5Padding(src, block.BlockSize())
	var dst []byte
	tmpData := make([]byte, block.BlockSize())
	for index := 0; index < len(paddingSrc); index += block.BlockSize() {
		block.Encrypt(tmpData, paddingSrc[index:index+block.BlockSize()])
		dst = append(dst, tmpData...)
	}
	return dst, nil
}

func (a *AesEcb) Decode(src []byte, secret []byte) ([]byte, error) {
	if len(src) == 0 {
		return src, nil // 空内容加密结果仍旧空
	}

	appendkeys := StringToBytes("秘钥长度仅支持16位、24位、32位，如果参数的秘钥有误，则按默认补足至32位方便使用")
	if !checkKey(secret) {
		secret = append(secret, appendkeys...)[:32]
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}
	if len(src)%block.BlockSize() != 0 {
		return nil, errors.New("源数据长度有误无法解密")
	}
	var dst []byte
	tmpData := make([]byte, block.BlockSize())

	for index := 0; index < len(src); index += block.BlockSize() {
		block.Decrypt(tmpData, src[index:index+block.BlockSize()])
		dst = append(dst, tmpData...)
	}
	dst, err = aesPpkcs5UnPadding(dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// PKCS5填充
func aesPkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去除PKCS5填充
func aesPpkcs5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])

	if length < unpadding {
		return nil, errors.New("invalid unpadding length")
	}
	return origData[:(length - unpadding)], nil
}

// 秘钥长度验证
func checkKey(key []byte) bool {
	n := len(key)
	return n == 16 || n == 24 || n == 32
}