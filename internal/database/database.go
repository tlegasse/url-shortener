package database

import (
	"database/sql"
	"os"
	"log"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Url struct {
	Id   int
	Time string
	Path string
	Url  string
}

type Db struct {
	Db *sql.DB
	errors map[string]error
}

func init() {
	c, err := sql.Open("sqlite3", "urls.db")

	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	Db := Db{
		Db: c,
		errors: map[string]error{
			"ErrDatabaseError": errors.New("Database error"),
			"ErrNoUrlFound": errors.New("No url found"),
		},
	}

	Db.SetupSchema("internal/database/schema.sql")
}

func (d Db) SetupSchema(f string) {
	// Read the sql file
	c, err := os.ReadFile(f)

	if err != nil {
		log.Fatalf("Failed to read the schema file: %v", err)
	}

	// Catch any errors
	_, err = d.Db.Exec(string(c))

	if err != nil {
		log.Fatalf("Failed to create the schema: %v", err)
	}
}

func (d Db) GetUrlFromPath(shortenedPath string) (Url, error) {
	var url Url
	rows, err := d.Db.Query("SELECT * FROM `urls` WHERE `path` = ?", shortenedPath)

	if err != nil {
		fmt.Println(err)
		return url, d.errors["ErrDatabaseError"]
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&url.Id, &url.Time, &url.Path, &url.Url)

		if err != nil {
			fmt.Println(err)
			return Url{},d.errors["ErrDatabaseError"]
		}
	} else {
		return url,d.errors["ErrNoUrlFound"]
	}

	return url, nil
}

func (d Db) InsertUrl(url Url) (error) {
	_, err := d.Db.Exec("INSERT INTO urls (path, url) VALUES (?, ?)", url.Path, url.Url)

	if err != nil {
		fmt.Println(err)
		return d.errors["ErrDatabaseError"]
	}

	return nil
}
