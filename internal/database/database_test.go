package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
    // Create a new mock SQL connection
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    // Call the Connect function with a mock filename
    testDB := Connect("test.db")

    // Assert that there were no errors when opening the connection
    assert.NotNil(t, testDB)

    // Ensure all expectations were met (if any were set)
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("There were unfulfilled expectations: %s", err)
    }
}

func TestSetupSchema(t *testing.T) {
    // Create a new mock SQL connection
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    // Assuming your schema file contains a command like "CREATE TABLE urls..."
    // Adjust the SQL command to match what's in your schema file
    mock.ExpectExec("CREATE TABLE IF NOT EXISTS urls").WillReturnResult(sqlmock.NewResult(0, 0))

    // Instantiate your DbType and set the mocked db instance
    d := DbType{
        instance: db,
    }

    // Call the SetupSchema function
    // You might need to adjust how the schema file is read or passed
    d.SetupSchema()

    // Assert all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("There were unfulfilled expectations: %s", err)
    }
}

func TestGetUrlFromPath(t *testing.T) {
    // Create a new mock SQL connection
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

	d := DbType{
		instance: db,
	}

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "time", "path", "url"}).
		AddRow(1, "" ,"1", "1")

	mock.ExpectQuery("^SELECT (.+) FROM urls WHERE path = (.+)$").WillReturnRows(rows)

	// Call the GetUrlFromPath function
	_, err = d.GetUrlFromPath("1")
	if err != nil {
		t.Errorf("Error in retrieving URL from path value: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertUrl(t *testing.T) {
    // Create a new mock SQL connection
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    d := DbType{
        instance: db,
    }

    u := Url{
        Path: "1",
        Url: "1",
    }

    // Use ExpectExec for INSERT operation, and use regular expression for the query
    mock.ExpectExec("^INSERT INTO urls").WithArgs(u.Path, u.Url).WillReturnResult(sqlmock.NewResult(1, 1))

    // Call the InsertUrl function
    err = d.InsertUrl(u)
    if err != nil {
        t.Errorf("Error in inserting URL: %s", err)
    }

    // Ensure all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("There were unfulfilled expectations: %s", err)
    }
}
