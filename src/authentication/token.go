package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func ValidateToken(r *http.Request) error {
	tokenString := extractTokenFromHeader(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Invalid token")
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Method sign unexpected : %v", token.Header["alg"])
	}
	return config.Secret, nil
}

func extractTokenFromHeader(r *http.Request) string {
	token := r.Header.Get("Authorization")
	tokenSplited := strings.Split(token, " ")

	if len(tokenSplited) == 2 {
		return tokenSplited[1]
	}
	return ""
}

type TokenStruct struct {
	Expiration int64  `json:"expiration,omitempty"`
	Token      string `json:"token,omitempty"`
	TypeToken  string `json:"type,omitempty"`
}
