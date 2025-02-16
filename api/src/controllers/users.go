package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a user"))
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
