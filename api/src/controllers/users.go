package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	databaseConnector, err := database.ToConnect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer databaseConnector.Close()

	userRepository := repository.NewRepositoryUserDatabase(databaseConnector)
	user.ID, err = userRepository.Create(user)

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	userNameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	databaseConnector, err := database.ToConnect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer databaseConnector.Close()

	userRepository := repository.NewRepositoryUserDatabase(databaseConnector)
	usersResult, err := userRepository.FindByNameOrNick(userNameOrNick)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, usersResult)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding users"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
