package config

import (
	"database/sql"
	"log"
	"os"
)

/*
	Initialize the DB function for database connectivity
*/

var DB *sql.DB

func InitDB() {
	dbfile := os.Getenv("DATABASE_FILE")

	var err error
	DB, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatalf("Unable to connect to database : %s", err)
	}

	log.Println("Database connection successful")
}
