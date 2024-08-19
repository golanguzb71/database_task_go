// Package tasks inserting data from file should work max=1s
package tasks

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func InsertDataFromFile(db *sql.DB) {
	insertBranches(db, "/var/lib/postgresql/100_branches.txt")
	insertUsers(db, "/var/lib/postgresql/5_k_users.txt")
	insertProducts(db, "/var/lib/postgresql/1mln_products.txt")
}

func insertBranches(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasSuffix(line, ";") {
			line = strings.TrimSuffix(line, ";")
		}
		if line == "" {
			continue
		}

		_, err := db.Exec("INSERT INTO branch(name) VALUES ($1)", line)
		if err != nil {
			panic(fmt.Sprintf("failed to insert branch: %v", err))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("error reading file: %v", err))
	}

	fmt.Println("Data of branches loaded successfully!")
}
func insertUsers(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("failed to open file"))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasSuffix(line, ";") {
			line = strings.TrimSuffix(line, ";")
		}
		if line == "" {
			continue
		}
		_, err := db.Exec("INSERT INTO users(name) values ($1)", line)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("error reading file:"))
	}
}

func insertProducts(db *sql.DB, filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		panic(fmt.Sprintf("file does not exist: %s", filePath))
	}

	copyCmd := fmt.Sprintf("COPY product (name, income_price, outcome_price) FROM '%s' WITH (FORMAT csv, DELIMITER ';')", filePath)

	_, err := db.Exec(copyCmd)
	if err != nil {
		panic(fmt.Sprintf("failed to execute COPY command: %v", err))
	}

	fmt.Println("Data of products loaded successfully!")
}

// TEST NOT PASSED
/*
/home/elon/.cache/JetBrains/GoLand2024.2/tmp/GoLand/___BenchmarkInsertDataFromFile_in_awesomeProject_PostgreSql_test.test -test.v -test.paniconexit0 -test.bench ^\QBenchmarkInsertDataFromFile\E$ -test.run ^$
goos: linux
goarch: amd64
pkg: awesomeProject/PostgreSql/test
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkInsertDataFromFile
Data of branches loaded successfully!
Data of products loaded successfully!
BenchmarkInsertDataFromFile-16    	       1	3296994467 ns/op
PASS
*/
