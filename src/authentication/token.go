package authentication

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	secretKey := []byte("Secret")

	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["usuarioId"] = userID
	fmt.Println(permissions)

	//Secret
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissions)
	fmt.Println(token)
	return token.SignedString(secretKey)
}
