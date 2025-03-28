package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	
	fmt.Println("Waiting for db startup ...")
	time.Sleep(5 * time.Second)

	db, err = sql.Open("postgres", connStr)
	if err != nil{
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatalf("Error connecting database: %v", err)
	}

	fmt.Println("successfully connected to database")
}

func GetDB() *sql.DB{
	return db
}

func CloseDB(){
	if err := db.Close(); err != nil{
		log.Fatal("Error closing database")
	}
}