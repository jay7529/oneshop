package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

var commonIV = Config("crypt.commonIV")
var key = Config("crypt.key")

// 加密字串
func Decrypt(text string) string {
	c, err := aes.NewCipher([]byte(key))
	CheckErr(err)
	cfb := cipher.NewCFBEncrypter(c, []byte(commonIV))
	plaintext := make([]byte, len([]byte(text)))
	cfb.XORKeyStream(plaintext, []byte(text))
	return string(plaintext)
}

// 解密字串
func Encrypt(text string) string {
	c, err := aes.NewCipher([]byte(key))
	CheckErr(err)
	cfb := cipher.NewCFBDecrypter(c, []byte(commonIV))
	plaintext := make([]byte, len([]byte(text)))
	cfb.XORKeyStream(plaintext, []byte(text))
	return string(plaintext)
}

// MD5加密
func MD5crypt(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
