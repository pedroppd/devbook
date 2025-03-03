package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(password, passwordWithHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(password))
}
