package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Prepare("register"); err != nil {
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
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
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
	usersResult, err := userRepository.FindByID(id)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, usersResult)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	userIDFromToken, err := authentication.GetUserIdFromToken(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	//Getting parameter
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userIDFromToken != id {
		fmt.Printf("User not allowed - %d", userIDFromToken)
		responses.Erro(w, http.StatusForbidden, err)
		return
	}

	//Getting body
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

	if err = user.Prepare("update"); err != nil {
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

	userRepository := repository.NewRepositoryUserDatabase(databaseConnector)
	err = userRepository.Update(id, user)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//Getting parameter
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userIDFromToken, err := authentication.GetUserIdFromToken(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userIDFromToken != id {
		errorMessage := fmt.Sprintf("User not allowed - %d", userIDFromToken)
		responses.Erro(w, http.StatusForbidden, errors.New(errorMessage))
		return
	}

	databaseConnector, err := database.ToConnect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer databaseConnector.Close()

	userRepository := repository.NewRepositoryUserDatabase(databaseConnector)
	if err := userRepository.DeleteByID(id); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	//Getting parameter
	vars := mux.Vars(r)
	userIDToFollow, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userID, err := authentication.GetUserIdFromToken(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userID == userIDToFollow {
		errorMessage := fmt.Sprintf("Cant follow youself - %d - %d", userID, userIDToFollow)
		responses.Erro(w, http.StatusForbidden, errors.New(errorMessage))
		return
	}

	databaseConnector, err := database.ToConnect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer databaseConnector.Close()

	userRepository := repository.NewRepositoryUserDatabase(databaseConnector)
	if err := userRepository.Follow(userID, userIDToFollow); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	//Getting parameter
	vars := mux.Vars(r)
	userIDToUnFollow, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userID, err := authentication.GetUserIdFromToken(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userID == userIDToUnFollow {
		errorMessage := fmt.Sprintf("Cant unfollow youself - %d - %d", userID, userIDToUnFollow)
		responses.Erro(w, http.StatusForbidden, errors.New(errorMessage))
		return
	}

	databaseConnector, err := database.ToConnect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer databaseConnector.Close()

	userRepository := repository.NewRepositoryUserDatabase(databaseConnector)
	if err := userRepository.UnFollow(userID, userIDToUnFollow); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func FindFollowers(w http.ResponseWriter, r *http.Request) {
	//Getting parameter
	vars := mux.Vars(r)
	userIDToFollow, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
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
	users, err := userRepository.FindFollowersByID(userIDToFollow)

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, users)
}

func FindFollowing(w http.ResponseWriter, r *http.Request) {
	//Getting parameter
	vars := mux.Vars(r)
	userIDFollowing, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
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
	users, err := userRepository.FindFollowingByID(userIDFollowing)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, users)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	//Getting parameter
	vars := mux.Vars(r)
	userIDFromPath, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	userID, err := authentication.GetUserIdFromToken(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userID != userIDFromPath {
		errorMessage := fmt.Sprintf("Cant update passoword of another user- %d - %d", userID, userIDFromPath)
		responses.Erro(w, http.StatusForbidden, errors.New(errorMessage))
		return
	}

	//Getting body
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var passoword models.Password

	if err = json.Unmarshal(requestBody, &passoword); err != nil {
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
	userPassword, err := userRepository.FindPassword(userIDFromPath)

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	checkErr := security.CheckPassword(userPassword, passoword.Current)
	if checkErr != nil {
		responses.Erro(w, http.StatusInternalServerError, errors.New("Password is not the current one"))
		return
	}

	hashNewPassword, errHashPassword := security.Hash(passoword.New)
	if errHashPassword != nil {
		responses.Erro(w, http.StatusInternalServerError, errors.New("Something went wrog to try create a hash for the password"))
		return
	}

	errUpdate := userRepository.UpdatePassword(userID, string(hashNewPassword))

	if errUpdate != nil {
		responses.Erro(w, http.StatusInternalServerError, errors.New("Something went wrog to try update the password"))
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
