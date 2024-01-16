// Description: Database interface for the url shortener
package database

import (
	"database/sql"
	"os"
	"log"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// URL struct for database operations
type Url struct {
	Id   int
	Path string // Path is the path segment corresponding to a stored URL
	Url  string
	Time string
}

type DbInterface interface {
	SetupSchema(f string)
	GetUrlFromPath(path string) (Url, error)
	InsertUrl(url Url) (error)
}

type DbType struct {
	instance *sql.DB
	errors map[string]error
}

// Database instance
var Db DbType
var dbFilename string = "urls.db"

func init() {
	Db.instance = Connect(dbFilename)

	Db.errors = map[string]error{
		"ErrDatabaseError": errors.New("Database error"),
		"ErrNoUrlFound": errors.New("No url found"),
	}

	Db.SetupSchema("./schema.sql")
}

func Connect(dbFilename string) *sql.DB {
	c, err := sql.Open("sqlite3", dbFilename)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	return c
}

func (d *DbType) SetupSchema(f string) {
	c, err := os.ReadFile(f)
	if err != nil {
		log.Fatalf("Failed to read the schema file: %v", err)
	}

	_, err = d.instance.Exec(string(c))
	if err != nil {
		log.Fatalf("Failed to create the schema: %v", err)
	}
}

// GetUrlFromPath returns a URL from the Database based on the path segment
func (d *DbType) GetUrlFromPath(path string) (Url, error) {
	var url Url

	rows, err := d.instance.Query("SELECT * FROM urls WHERE path = ?", path)
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

func (d *DbType) InsertUrl(url Url) (error) {
	_, err := d.instance.Exec("INSERT INTO urls (path, url) VALUES (?, ?)", url.Path, url.Url)
	if err != nil {
		fmt.Println(err)
		return d.errors["ErrDatabaseError"]
	}

	return nil
}
