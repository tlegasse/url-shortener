package urlshortener

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"math/rand"

	database "github.com/tlegasse/url-shortener/internal/database"
)

type Shortener struct {
	db *sql.DB
}

var ShortenerInstance Shortener
var once sync.Once

func init() {
	once.Do(func() {
		ShortenerInstance = Shortener{
			db: database.DbInstance,
		}
	})
}

func (s *Shortener) Redirect(w http.ResponseWriter, r *http.Request) {
	// Get the path from the Request
	p := r.URL.Path[1:]

	// Get the URL from the database
	url, err := database.GetUrlFromPath(p)

	if err != nil {
		fmt.Println(err)
		// Write a response to the page that reports a 404 to the user with a short message and an error code of 404
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}

	// Redirect the user to the URL
	http.Redirect(w, r, url.Url, http.StatusSeeOther)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func (s *Shortener) Shorten(w http.ResponseWriter, r *http.Request) {
	// Get URL from request GET parameters
	u := r.URL.Query().Get("url")
	p := RandStringRunes(10)

	// Create a database.Url struct with values from the request
	url := database.Url{
		Path: p,
		Url:  u,
	}

	// Insert the url into the database
	err := database.InsertUrl(url)

	if err != nil {
		fmt.Println(err)
	}

	// Write the new URL path to the Response
	fmt.Fprintf(w, "http://localhost:8080/%s", p)
}
