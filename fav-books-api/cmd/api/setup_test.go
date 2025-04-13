package main

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
