package SecureUtils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/wjdsg0327/wjdsgtools/RandomUtils"
)

type SecureUtils struct {
}

// 参数：需要加密的字符串
func getMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5 md5加密（无法解密）
func (SecureUtils) Md5(str string, salt string) string {
	return getMd5(getMd5(str + salt))
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// Encrypt 加密
/**
 * @params text:需要加密的数据
 * @params key :自定义加密密钥，如果不需要自定义，则传空字符串，系统自动生成密码并且返回
 * 注意：自定义密钥必须等于32比特
 * return：加密之后内容，密钥
 */
func (SecureUtils) Encrypt(text string, key string) (string, string, error) {

	var code []byte
	if key == "" {
		randomString := RandomUtils.RandomUtils{}.GenerateRandomString(true, true, false, true, 32)
		code = []byte(randomString)
	} else if len(key) != 32 {
		return "", "", errors.New(fmt.Sprintf("密码长度不够，需要长度32，您的长度：%d", len(key)))

	} else {
		code = []byte(key)
	}
	c, err := aes.NewCipher(code)
	if err != nil {
		return "", "", err
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(text))
	cfb.XORKeyStream(ciphertext, []byte(text))

	return base64.StdEncoding.EncodeToString(ciphertext), string(code), nil
}

// Decrypt 解密函数
// 参数：ciphertext加密之后的内容，key密码
// 返回值：解密内容
func (SecureUtils) Decrypt(ciphertext string, key string) (string, error) {

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
