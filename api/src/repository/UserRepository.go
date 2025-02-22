package repository

import (
	"api/src/models"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewRepositoryUserDatabase(db *sql.DB) *User {
	return &User{db}
}

//Create - Create a user in database
func (repository User) Create(user models.User) (uint64, error) {
	stmt, err := repository.db.Prepare("insert into users (username, nick, email, userpassword) values (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	resultset, err := stmt.Exec(user.Nome, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := resultset.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}
