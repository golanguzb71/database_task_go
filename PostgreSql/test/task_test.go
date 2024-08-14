package test

import (
	"awesomeProject/PostgreSql/tasks"
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
)

func BenchmarkInsertDataFromFile(b *testing.B) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		b.Fatalf("Failed to connect database")
	}
	defer db.Close()
	for i := 0; i < b.N; i++ {
		b.ResetTimer()
		tasks.InsertDataFromFile(db)
	}
}
