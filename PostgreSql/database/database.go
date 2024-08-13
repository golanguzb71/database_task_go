package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func ConnectPostgres() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		"postgres",
		"postgres",
		"postgres",
		"localhost",
		"5432",
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	fmt.Println("Database connection successful")
}
