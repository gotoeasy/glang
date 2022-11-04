package cmn

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 使用公钥进行RSA加密后按Base64编码字符串
func EncodeRsa(str string, pubKeyFileName string) (string, error) {
	bt, err := EncodeRsaBytes(StringToBytes(str), pubKeyFileName)
	if err != nil {
		return "", err
	}
	return Base64(bt), nil
}

// 按Base64解码字符串后使用私钥进行RSA解密
func DecodeRsa(str string, pubKeyFileName string) (string, error) {
	by, err := Base64Decode(str)
	if err != nil {
		return "", err
	}
	bt, err := DecodeRsaBytes(by, pubKeyFileName)
	if err != nil {
		return "", err
	}
	return BytesToString(bt), nil
}

// 当前目录下创建4096位的秘钥文件"rsa_private.pem、rsa_public.pem"
func GenerateRsaKey() error {
	return GenerateRsaKeyFile(2048, "rsa_private.pem", "rsa_public.pem")
}

// 使用公钥进行RSA加密
func EncodeRsaBytes(data []byte, pubKeyFileName string) ([]byte, error) {
	file, err := os.Open(pubKeyFileName)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	cipherTextBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
	if err != nil {
		return nil, err
	}
	return cipherTextBytes, nil
}

// 使用私钥进行RSA解密
func DecodeRsaBytes(data []byte, privateKeyFileName string) ([]byte, error) {
	file, err := os.Open(privateKeyFileName)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, fileInfo.Size())
	defer file.Close()
	file.Read(buf)

	block, _ := pem.Decode(buf)

	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priKey, data)
}

// 创建秘钥文件（keySize通常是1024、2048、4096）
func GenerateRsaKeyFile(keySize int, priKeyFile, pubKeyFile string) error {
	// private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return err
	}

	derText := x509.MarshalPKCS1PrivateKey(privateKey)

	block := pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}

	file, err := os.Create(priKeyFile)
	if err != nil {
		return err
	}

	pem.Encode(file, &block)
	file.Close()

	// public key
	publicKey := privateKey.PublicKey

	derpText, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derpText,
	}

	//file,err = os.Create("rsa_public.pem")
	file, err = os.Create(pubKeyFile)
	if err != nil {
		return err
	}
	pem.Encode(file, &block)
	file.Close()

	return nil
}
