package urlshortener

import (
	"fmt"
	"net/http"
	"math/rand"

	database "github.com/tlegasse/url-shortener/internal/database"
)

type ShortenerType struct {
	Url string
	Port string
}

var Shortener ShortenerType

func Setup(url string, port string) {
	Shortener.Url = url
	Shortener.Port = port

	Shortener.SetupRoutes()
}

func (s ShortenerType) SetupRoutes() {
	http.HandleFunc("/shorten", s.Shorten)
	http.HandleFunc("/", s.Redirect)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func CleanUrl(u string) string {
	if u[len(u)-1:] == "/" {
		return u[:len(u)-1]
	}
	return u
}

func (s ShortenerType) Shorten(w http.ResponseWriter, r *http.Request) {
	// Get URL from request GET parameters
	u := r.URL.Query().Get("url")
	p := RandStringRunes(10)

	// Create a database.Url struct with values from the request
	url := database.Url{
		Path: p,
		Url:  u,
	}

	// Insert the url into the database
	err := database.Db.InsertUrl(url)
	if err != nil {
		fmt.Println(err)
	}

	baseUrl := CleanUrl(s.Url)

	// Write the new URL path to the Response
	_, err = w.Write([]byte(baseUrl + ":" + s.Port + "/" + p))
	if err != nil {
		fmt.Println(err)
	}
}

func (s ShortenerType) Redirect(w http.ResponseWriter, r *http.Request) {
	// Get the path from the Request
	p := r.URL.Path[1:]

	// Get the URL from the database
	url, err := database.Db.GetUrlFromPath(p)

	if err != nil {
		fmt.Println(err)
		// Write a response to the page that reports a 404 to the user with a short message and an error code of 404
		http.Error(w, "404 Not Found", http.StatusNotFound)
	} else {
		// Redirect the user to the URL
		http.Redirect(w, r, url.Url, http.StatusSeeOther)
	}
}

