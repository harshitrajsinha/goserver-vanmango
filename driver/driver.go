package driver

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

// Initialize database connection
func InitDB() error {

	// setup connection url
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("Waiting for db startup ...")
	time.Sleep(5 * time.Second) // wait for 5 seconds to complete database setup

	// open database connection
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// connection will be closed in main.go

	// verify connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error during database Ping(): %v", err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("error closing database connection")
	}
	return nil
}
