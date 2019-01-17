package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var blockExplorerDb *sqlx.DB

const (
	defaultDb      = "shyftdb"
	connStrDocker  = "user=postgres host=pg password=docker sslmode=disable"
	connStrDefault = "user=postgres host=localhost sslmode=disable"
)

// SPGDatabase struct return db type
type SPGDatabase struct {
	Db *sqlx.DB // PostgresDB instance
}

// ConnectionStr - return a Connection to the PG admin database
func ConnectionStr() string {
	dbEnv := os.Getenv("DBENV")
	switch dbEnv {
	default:
		return connStrDefault
	case "docker":
		return connStrDocker
	}
}

// ShyftConnectStr - Returns the Connection String With The appropriate database
func ShyftConnectStr() string {
	return fmt.Sprintf("%s%s", ConnectionStr(), " dbname=shyftdb")
}

// Connect - return a connection to a postgres database wi
func Connect(connectURL string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", connectURL)
	if err != nil {
		fmt.Println("ERROR OPENING DB, NOT INITIALIZING")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// NewShyftDatabase returns a PostgresDB wrapped object.
func NewShyftDatabase() (*SPGDatabase, error) {
	if blockExplorerDb == nil {
		log.Printf("DB Connected !!")
		blockExplorerDb = Connect(ShyftConnectStr())
		conn := blockExplorerDb
		conn.Ping()
		return &SPGDatabase{
			Db: conn,
		}, nil
	}
	conn := blockExplorerDb
	conn.Ping()
	return &SPGDatabase{
		Db: conn,
	}, nil
}

// ReturnShyftDatabase returns a PostgresDB wrapped object.
func ReturnShyftDatabase() (*SPGDatabase, error) {
	return NewShyftDatabase()
}
