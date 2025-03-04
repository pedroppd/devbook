package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	//Database connection
	databaseConnector, err := database.ToConnect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer databaseConnector.Close()

	//Repository
	userRepository := repository.NewRepositoryUserDatabase(databaseConnector)
	user.ID, err = userRepository.Create(user)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

}
