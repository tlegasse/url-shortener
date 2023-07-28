package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const tableName string = "urls"
const file string = "urls.db"
const create string = `CREATE TABLE IF NOT EXISTS [` + tableName + `] (
	id INTEGER NOT NULL PRIMARY KEY,
	time DATETIME  NOT NULL,
	path TEXT NOT NULL,
	url TEXT NOT NULL
);`

func Setup() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return db, nil
}
