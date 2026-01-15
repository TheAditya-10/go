package db

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMySQL(dsn string) {
	var err error
	
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}
	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging MySQL:", err)
	}
	log.Println("Connected to MySQL database successfully")
}