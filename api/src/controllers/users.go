package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	databaseConnector, err := database.ToConnect()
	if err != nil {
		log.Fatal(err)
	}
	print("Entrei no método create user userRepository")
	userRepository := repository.NewRepositoryUserDatabase(databaseConnector)
	userID, err := userRepository.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("Id inserted successfully : %d", userID)))
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding users"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
