package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"errors"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.Erro(w, http.StatusBadRequest, errors.New("invalid token"))
			return
		}

		next(w, r)
	}
}
