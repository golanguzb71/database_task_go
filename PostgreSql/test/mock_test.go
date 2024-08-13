package test

import (
	"awesomeProject/PostgreSql/mock"
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
)

func BenchmarkInsertBranches(b *testing.B) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM branch")
	if err != nil {
		b.Fatalf("Failed to clean database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mock.InsertBranches(db)
	}
}

// TEST PASSED
// 100 TPS passed 100 ms

func BenchmarkInsertUsers(b *testing.B) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		b.Fatalf("Failed to clean database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mock.InsertUsers(db)
	}
}

func BenchmarkInsertProduct(b *testing.B) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM product")
	if err != nil {
		b.Fatalf("Failed to clean database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mock.InsertProducts(db)
	}
}

// TEST PASSED
// 100 TPS passed 100 ms

func BenchmarkInsertOrders(b *testing.B) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM orders")
	if err != nil {
		b.Fatalf("Failed to clean database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mock.InsertOrders(db)
	}
}

func BenchmarkInsertOrderItems(b *testing.B) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM order_items")
	if err != nil {
		b.Fatalf("Failed to clean database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mock.InsertOrderItems(db)
	}
}

// TEST PASSED
// 100 TPS passed 100 ms

/*

goos: linux
goarch: amd64
pkg: awesomeProject/PostgreSql/test
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkInsertBranches
Inserted 10,000 branches!
BenchmarkInsertBranches-16      	       1	7251997113 ns/op
BenchmarkInsertUsers
Inserted 10,000 users!
BenchmarkInsertUsers-16         	       1	8049150205 ns/op
BenchmarkInsertProduct
Inserted 10,000 products!
BenchmarkInsertProduct-16       	       1	8214362562 ns/op
BenchmarkInsertOrders
Inserted 10,000 orders!
BenchmarkInsertOrders-16        	       1	8084874015 ns/op
BenchmarkInsertOrderItems
Inserted 10,000 order items!
BenchmarkInsertOrderItems-16    	       1	8196023723 ns/op
PASS
*/
