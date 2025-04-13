package main

/* Data-Dog sql mock vs docker database
Hi Trevor, I saw in other courses of yours that you use a docker db for the test database,
it does have a bit more of a setup but I'm guessing it stays true to the original db behavior.
This method does look a lot faster to setup.

Trevor: I always use Docker during development now.
*/

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kons77/fav-books-api/internal/data"
)

var testApp application
var mockedDB sqlmock.Sqlmock

func TestMain(m *testing.M) {
	testDB, myMock, _ := sqlmock.New()
	mockedDB = myMock

	defer testDB.Close()

	testApp = application{
		config:      config{},
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog:    log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime),
		models:      data.New(testDB),
		environment: "development",
	}

	os.Exit(m.Run())
}
