package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
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
	userResponse, err := userRepository.FindByEmail(user.Email)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(userResponse.Password, user.Password); err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	token, err := authentication.CreateToken(userResponse.ID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
