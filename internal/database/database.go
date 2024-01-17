// Description: Database interface for the url shortener
package database

import (
	"database/sql"
	"log"
	"errors"
	"fmt"
    "embed"

	_ "github.com/mattn/go-sqlite3"
	util "github.com/tlegasse/url-shortener/internal/util"
)

//go:embed schema.sql
var schemaFS embed.FS

// URL struct for database operations
type Url struct {
	Id   int
	Path string // Path is the path segment corresponding to a stored URL
	Url  string
	Time string
}

type DbInterface interface {
	SetDatabase(f string)
	SetupSchema(f string)
	GetUrlFromPath(path string) (Url, error)
	InsertUrl(url Url) (error)
}

type DbType struct {
	Instance *sql.DB
	Errors map[string]error
}

// Database Instance
var Db DbType

func init() {
	c := util.GetConfig()

	Db.SetDatabase(Connect(c.DbPath))

	Db.Errors = map[string]error{
		"ErrDatabaseError": errors.New("Database error"),
		"ErrNoUrlFound": errors.New("No url found"),
	}

	Db.SetupSchema()
}

func (d *DbType) SetDatabase(db *sql.DB) {
	d.Instance = db
}


func Connect(dbFilename string) *sql.DB {
	c, err := sql.Open("sqlite3", dbFilename)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	return c
}

func (d *DbType) SetupSchema() {
    content, err := schemaFS.ReadFile("schema.sql")
    if err != nil {
        log.Fatalf("Failed to read the schema file: %v", err)
    }

	_, err = d.Instance.Exec(string(content))
	if err != nil {
		log.Fatalf("Failed to create the schema: %v", err)
	}
}

// GetUrlFromPath returns a URL from the Database based on the path segment
func (d *DbType) GetUrlFromPath(path string) (Url, error) {
	var url Url

	rows, err := d.Instance.Query("SELECT * FROM urls WHERE path = ?", path)
	if err != nil {
		fmt.Println(err)
		return url, d.Errors["ErrDatabaseError"]
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&url.Id, &url.Time, &url.Path, &url.Url)
		fmt.Printf("Url: %v\n", url)
		if err != nil {
			fmt.Println(err)
			return Url{},d.Errors["ErrDatabaseError"]
		}
	} else {
		return url,d.Errors["ErrNoUrlFound"]
	}

	return url, nil
}

func (d *DbType) InsertUrl(url Url) (error) {
	_, err := d.Instance.Exec("INSERT INTO urls (path, url) VALUES (?, ?)", url.Path, url.Url)
	if err != nil {
		fmt.Println(err)
		return d.Errors["ErrDatabaseError"]
	}

	return nil
}
