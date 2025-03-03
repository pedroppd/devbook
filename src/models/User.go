package models

import (
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

func (user *User) removeWhiteSpace() {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}
	user.removeWhiteSpace()
	return nil
}
