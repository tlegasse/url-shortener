package urlshortener

import (
	"testing"
	"log"
	"net/http/httptest"
	"strings"

	database "github.com/tlegasse/url-shortener/internal/database"

	"github.com/DATA-DOG/go-sqlmock"
)

var shortener ShortenerType

const (
	hostname = "http://localhost"
	port = "8080"
)

func TestMain(m *testing.M) {

	Setup(hostname, port, database.Db)

	shortener = Shortener

	m.Run()
}

func TestSetup(t *testing.T) {
	// Test if the Shortener contains the hostname and port as specified
	if shortener.Url != hostname {
		t.Errorf("Shortener.Hostname is not initialized")
	}

	if shortener.Port != port {
		t.Errorf("Shortener.Port is not initialized")
	}
}

func TestRandStringRunes(t *testing.T) {
	str1 := RandStringRunes(10)
	str2 := RandStringRunes(10)

	if str1 == str2 {
		t.Errorf("RandStringRunes(10) returned the same string twice")
	}
}

func TestRandStringRunesLen(t *testing.T) {
	for i := 0; i < 100; i++ {
		str := RandStringRunes(i)
		if len(str) != i {
			t.Errorf("RandStringRunes(%d) returned a string with length %d", i, len(str))
		}
	}
}


func TestCleanUrl(t *testing.T) {
	dirtyUrl := "http://www.example.com/"
	cleanUrl := CleanUrl(dirtyUrl)

	if cleanUrl != "http://www.example.com" {
		t.Errorf("CleanUrl(%s) returned %s", dirtyUrl, cleanUrl)
	}
}

func TestShorten(t *testing.T) {
    db, _, err := sqlmock.New()
    if err != nil {
		log.Fatalf("Error connecting to mock database: %v", err)
    }
    defer db.Close()

	shortener.Db.SetDatabase(db)

	r := httptest.NewRequest("GET", "/shorten?url=http://localhost", nil)
	w := httptest.NewRecorder()

	shortener.Shorten(w, r)

	if w.Code != 200 {
		t.Errorf("Shorten returned %d", w.Code)
	}

	if !strings.HasPrefix(w.Body.String(), hostname + ":" + port + "/") {
		t.Errorf("Shorten returned %s", w.Body.String())
	}
}

func TestRedirect(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
		log.Fatalf("Error connecting to mock database: %v", err)
    }
    defer db.Close()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/path1", nil)

	shortener.Db.SetDatabase(db)

	rows := sqlmock.NewRows([]string{"id", "path", "url"}).AddRow(1, "path1", "http://example1.com")

	mock.ExpectQuery("^SELECT (.+) FROM urls WHERE path = (.+)$").WillReturnRows(rows)

	shortener.Redirect(w, r)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if w.Code != 302 {
		t.Errorf("Shorten returned %d", w.Code)
	}

	if !strings.HasPrefix(w.Body.String(), hostname + ":" + port + "/") {
		t.Errorf("Shorten returned %s", w.Body.String())
	}
}
