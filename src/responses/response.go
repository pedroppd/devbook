package responses

import (
	"encoding/json"
	"errors"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			// Apenas registra o erro, sem tentar escrever outra resposta
			Erro(w, http.StatusInternalServerError, errors.New(`{"erro": "Erro ao serializar resposta"}`))
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	// Garante que a função de erro não chame JSON de novo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resposta := struct {
		Error string `json:"erro"`
	}{
		Error: erro.Error(),
	}

	json.NewEncoder(w).Encode(resposta) // Apenas escreve a resposta de erro
}
