package authentication

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["usuarioId"] = userID
	//Secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.Secret)
}
