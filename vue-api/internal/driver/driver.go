package driver

import (
	"database/sql"
	"fmt"
	"time"

	// _ "github.com/jackc/pgx/pgconn" - now it's a part of v5
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpedDbConn = 5
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectPosgres creates database pool for postgres
func ConnectPosgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpedDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = d

	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error!", err)
		return err
	}

	return nil
}
