package encrypts

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Md5(str string) string {
	hash := md5.New()
	io.WriteString(hash, str)
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha256(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)
	return hex.EncodeToString(hash.Sum(nil))
}
