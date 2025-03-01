package models

import (
	"errors"
	"strings"
)

type User struct {
	ID        uint64 `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

func (user *User) validate() error {
	if user.Nome == "" {
		return errors.New("Name cannot be null")
	}

	if user.Nick == "" {
		return errors.New("Name cannot be null")
	}

	if user.Email == "" {
		return errors.New("Name cannot be null")
	}

	if user.Password == "" {
		return errors.New("Name cannot be null")
	}

	return nil
}

func (user *User) removeWhiteSpace() {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *User) Prepare() error {
	if erro := user.validate(); erro != nil {
		return erro
	}
	user.removeWhiteSpace()
	return nil
}
