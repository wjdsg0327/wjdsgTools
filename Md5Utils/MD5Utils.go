package Md5Utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/wjdsg0327/wjdsgtools/randomUtils"
)

type Md5Utils struct {
}

// 参数：需要加密的字符串
func getMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5 md5加密（无法解密）
func (Md5Utils) Md5(str string, salt string) string {
	return getMd5(getMd5(str + salt))
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// Encrypt 加密
// 参数1 text:需要加密的数据
// 返回值：加密之后内容，密钥
func (Md5Utils) Encrypt(text string) (string, string, error) {

	randomString := randomUtils.RandomUtils{}.GenerateRandomString(true, true, false, true, 32)

	key := []byte(randomString)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(text))
	cfb.XORKeyStream(ciphertext, []byte(text))

	return base64.StdEncoding.EncodeToString(ciphertext), string(key), nil
}

// Decrypt 解密函数
// 参数：ciphertext加密之后的内容，key密码
// 返回值：解密内容
func (Md5Utils) Decrypt(ciphertext string, key string) (string, error) {

	byteSlice := []byte(key)

	// 解码Base64
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// 创建AES解密算法
	c, err := aes.NewCipher(byteSlice)
	if err != nil {
		return "", err
	}

	// 使用AES解密算法和初始化向量创建CFB模式解密器
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintext := make([]byte, len(decodedCiphertext))
	cfbdec.XORKeyStream(plaintext, decodedCiphertext)

	return string(plaintext), nil
}
