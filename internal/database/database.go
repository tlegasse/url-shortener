package database

import (
	"database/sql"
	"os"
	"log"
	"sync"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrDatabaseError   = errors.New("Database error")
	ErrNoUrlFound      = errors.New("No url found")
)

var DbInstance *sql.DB
var once sync.Once

type Url struct {
	Id   int
	Time string
	Path string
	Url  string
}

func init() {
	CreateConnection("urls.db")
	SetupSchema("internal/database/schema.sql")
}

func SetupSchema(f string) {
	// Read the sql file
	c, err := os.ReadFile(f)

	if err != nil {
		log.Fatalf("Failed to read the schema file: %v", err)
	}

	// Catch any errors
	_, err = DbInstance.Exec(string(c))

	if err != nil {
		log.Fatalf("Failed to create the schema: %v", err)
	}
}

func CreateConnection(dbFile string) {
	once.Do(func() {
		var err error

		DbInstance, err = sql.Open("sqlite3", dbFile)

		if err != nil {
			log.Fatalf("Failed to open the database: %v", err)
		}
	})
}

func GetUrlFromPath(shortenedPath string) (Url, error) {
	var url Url
	rows, err := DbInstance.Query("SELECT * FROM `urls` WHERE `path` = ?", shortenedPath)

	if err != nil {
		fmt.Println(err)
		return url, ErrDatabaseError
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&url.Id, &url.Time, &url.Path, &url.Url)

		if err != nil {
			fmt.Println(err)
			return Url{}, ErrDatabaseError
		}
	} else {
		return url, ErrNoUrlFound
	}

	return url, nil
}

func InsertUrl(url Url) (error) {
	_, err := DbInstance.Exec("INSERT INTO urls (path, url) VALUES (?, ?)", url.Path, url.Url)

	if err != nil {
		fmt.Println(err)
		return ErrDatabaseError
	}

	return nil
}
