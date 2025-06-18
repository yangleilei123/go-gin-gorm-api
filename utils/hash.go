package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 将明文密码加密为 bcrypt 哈希
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
