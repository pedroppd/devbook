package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type User struct {
	db *sql.DB
}

func NewRepositoryUserDatabase(db *sql.DB) *User {
	return &User{db}
}

// Create - Create a user in database
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

// Create - Create a user in database
func (repository User) Update(userID uint64, user models.User) error {
	stmt, err := repository.db.Prepare("update users set username = ?, nick = ?, email = ? where id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Nome, user.Nick, user.Email, userID)
	if err != nil {
		return err
	}
	return nil
}

func (repository User) FindByID(id uint64) (models.User, error) {
	var user models.User

	statement, erro := repository.db.Prepare("select id, username, nick, email, createdAt from users where id = ?")

	if erro != nil {
		return user, erro
	}

	defer statement.Close()

	err := statement.QueryRow(id).Scan(&user.ID, &user.Nome, &user.Nick, &user.Email, &user.CreatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository User) FindByNameOrNick(nameOrNick string) ([]models.User, error) {
	queryParam := fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repository.db.Query("select id, username, nick, email, createdAt from users where username like ? or nick like ?", queryParam, queryParam)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if erro = lines.Scan(&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CreatedAt); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}
