package encrypts

import (
	"fmt"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	plain := "123456"
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "abcdefgehjhijkmlkjjwwoew"
	text, err := Encrypt(plain, key)
	fmt.Println(text, err)
	// 解密
	deText, err := Decrypt(text, key)
	fmt.Println("-----------------------")
	fmt.Println(deText, err)
}
