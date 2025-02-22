package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ToConnect() (*sql.DB, error) {
	fmt.Println(config.ConnectionDatabase)
	db, err := sql.Open("mysql", config.ConnectionDatabase)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
