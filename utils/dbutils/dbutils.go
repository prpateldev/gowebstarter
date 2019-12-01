package dbutils

import "gowebstarter/configs/dbconfig"

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

// Return connected database instance
func GetDB() *sql.DB {
	return database
}

// Set database to package level database
func SetDB(db *sql.DB) {
	database = db
}

// Connect to database
func Connect(config dbconfig.DatabaseConfig) (*sql.DB, error) {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DBUser, config.DBPass,
		config.DBHost, config.DBDbase)
	database, err := sql.Open("mysql", dbConn)

	return database, err
}

// To connect and initiate package database object and panic on error
func Init(config dbconfig.DatabaseConfig) {
	db, err := Connect(config)
	if err != nil {
		log.Println("Couldn't connect!", err.Error())
		panic(err)
	} else {
		SetDB(db)
		log.Println("Database connected on ", config.DBHost+config.DBPort)
	}
}
