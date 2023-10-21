package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256V 函数用于计算字符串的 SHA-256 哈希值
func SHA256V(str string) string {
	message := []byte(str)
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
