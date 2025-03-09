package responses

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			Erro(w, http.StatusInternalServerError, erro)
		}
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Error string `json:"erro"`
	}{
		Error: erro.Error(),
	})
}
