package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 将明文密码生成 bcrypt 哈希
func HashPassword(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ComparePassword 比较明文密码与哈希是否匹配
func ComparePassword(hashed, plain string) bool {
	fmt.Printf("Comparing hashed: %s with plain: %s\n\n", hashed, plain) // Debugging line
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
