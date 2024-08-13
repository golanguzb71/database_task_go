package main

import (
	"awesomeProject/PostgreSql/database"
	"awesomeProject/PostgreSql/mock"
)

func main() {
	database.ConnectPostgres()
	db := database.DB
	mock.CreateTables(db)

}
