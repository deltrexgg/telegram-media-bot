package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

/*
	Initialize the DB function for database connectivity
*/


func InitDB() (*sql.DB, error){
	dbfile := os.Getenv("DATABASE_FILE")
	log.Println(dbfile)
	return sql.Open("sqlite3", dbfile)
	// if err != nil {
	// 	log.Fatalf("Unable to connect to database : %s", err)
	// }

	//log.Println("Database connection successful")
}
