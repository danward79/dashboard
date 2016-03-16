package models

import (
	"database/sql"

	//init the so-sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

//DB ..
type DB struct {
	*sql.DB
}

//NewDB create a new DB connection
func NewDB(dbPath string) (*DB, error) {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
