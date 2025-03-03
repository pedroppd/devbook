package models

import (
	"api/src/security"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64 `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

func (user *User) validate(step string) error {

	if user.Nome == "" {
		return errors.New("Name cannot be null")
	}

	if user.Nick == "" {
		return errors.New("Nick cannot be null")
	}

	if user.Email == "" {
		return errors.New("Email cannot be null")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("Email is invalid")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Password cannot be null")
	}

	return nil
}

func (user *User) Format(step string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if step == "register" {
		passowordHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passowordHash)
	}
	return nil
}

func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}
	if erro := user.Format(step); erro != nil {
		return erro
	}
	return nil
}
