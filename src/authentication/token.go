package authentication

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (TokenStruct, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["usuarioId"] = userID
	//Secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	tokenstring, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return TokenStruct{}, err
	}
	return TokenStruct{permissions["exp"].(int64), tokenstring, "Bearer"}, nil
}

type TokenStruct struct {
	Expiration int64  `json:"expiration,omitempty"`
	Token      string `json:"token,omitempty"`
	TypeToken  string `json:"type,omitempty"`
}
