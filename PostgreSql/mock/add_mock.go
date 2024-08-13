package mock

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func InsertMockData(db *sql.DB) {
	InsertBranches(db)
	InsertUsers(db)
	InsertOrders(db)
	InsertProducts(db)
	InsertOrderItems(db)
}

func InsertBranches(db *sql.DB) {
	for i := 0; i < 10000; i++ {
		name := fmt.Sprintf("Branch-%d", i+1)
		_, err := db.Exec("INSERT INTO branch (name) VALUES ($1)", name)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Inserted 10,000 branches!")
}

func InsertUsers(db *sql.DB) {
	for i := 0; i < 10000; i++ {
		name := fmt.Sprintf("User-%d", i+1)
		_, err := db.Exec("INSERT INTO users (name) VALUES ($1)", name)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Inserted 10,000 users!")
}

func InsertProducts(db *sql.DB) {
	for i := 0; i < 10000; i++ {
		name := fmt.Sprintf("Product-%d", i+1)
		incomePrice := rand.Float64() * 100
		outcomePrice := rand.Float64() * 150
		_, err := db.Exec("INSERT INTO product (name, income_price, outcome_price) VALUES ($1, $2, $3)", name, incomePrice, outcomePrice)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Inserted 10,000 products!")
}

func InsertOrders(db *sql.DB) {
	for i := 0; i < 10000; i++ {
		userID := rand.Intn(10000) + 1
		deliveredAt := time.Now().Add(time.Duration(-rand.Intn(10000)) * time.Hour)
		_, err := db.Exec("INSERT INTO orders (user_id, delivered_at) VALUES ($1, $2)", userID, deliveredAt)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Inserted 10,000 orders!")
}

func InsertOrderItems(db *sql.DB) {
	for i := 0; i < 10000; i++ {
		orderID := rand.Intn(10000) + 1
		productID := rand.Intn(10000) + 1
		_, err := db.Exec("INSERT INTO order_items (order_id, product_id) VALUES ($1, $2)", orderID, productID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Inserted 10,000 order items!")
}
