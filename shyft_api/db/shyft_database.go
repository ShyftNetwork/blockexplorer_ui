package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
)

var blockExplorerDb *sqlx.DB

const (
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
		logger.Log("DB Connected !!")
		blockExplorerDb = Connect(ShyftConnectStr())
		conn := blockExplorerDb
		if err := conn.Ping(); err != nil {
			logger.Warn("Database error: " + err.Error())
		}
		return &SPGDatabase{
			Db: conn,
		}, nil
	}
	conn := blockExplorerDb
	if err := conn.Ping(); err != nil {
		logger.Warn("Database error: " + err.Error())
	}
	return &SPGDatabase{
		Db: conn,
	}, nil
}

// ConnectShyftDatabase returns database of type *SPGDatabase
func ConnectShyftDatabase() *SPGDatabase {
	db, err := NewShyftDatabase()
	if err != nil {
		logger.Warn("Unable to connect to database: " + err.Error())
		return nil
	}
	return db
}