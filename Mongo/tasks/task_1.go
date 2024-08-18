package tasks

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strings"
	"sync"
)

func InsertingBranchesByIteratingFile(db *mongo.Database, fileUri string) {
	file, err := os.Open(fileUri)
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()
	collection := db.Collection("branches")
	scanner := bufio.NewScanner(file)
	var operations []mongo.WriteModel
	for scanner.Scan() {
		branchName := strings.TrimSpace(scanner.Text())
		if branchName == "" {
			continue
		}
		branchDoc := mongo.NewInsertOneModel().SetDocument(map[string]string{
			"name": branchName,
		})
		operations = append(operations, branchDoc)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error reading %v", err)
	}
	_, err = collection.BulkWrite(context.Background(), operations, options.BulkWrite())
	if err != nil {
		log.Fatalf("Failed to insert")
	}
	fmt.Println("successfully added")
}

func InsertingProductByCopyingFile(db *mongo.Database, fileUri string) {
	file, err := os.Open(fileUri)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", fileUri, err)
	}
	defer file.Close()

	collection := db.Collection("products")
	scanner := bufio.NewScanner(file)
	var operations []mongo.WriteModel
	batchSize := 20000
	var wg sync.WaitGroup
	sem := make(chan struct{}, 100)

	for scanner.Scan() {
		if len(operations) >= batchSize {
			wg.Add(1)
			sem <- struct{}{}
			go func(batch []mongo.WriteModel) {
				defer wg.Done()
				_, err := collection.BulkWrite(context.Background(), batch)
				if err != nil {
					log.Printf("Failed to insert batch: %v", err)
				}
				<-sem
			}(operations)
			operations = nil
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "product_name") {
			continue
		}
		prod := strings.Split(line, ";")
		if len(prod) != 3 {
			continue
		}
		prodDoc := mongo.NewInsertOneModel().SetDocument(map[string]interface{}{
			"product_name":  prod[0],
			"income_price":  prod[1],
			"outcome_price": prod[2],
		})
		operations = append(operations, prodDoc)
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	if len(operations) > 0 {
		wg.Add(1)
		go func(batch []mongo.WriteModel) {
			defer wg.Done()
			_, err := collection.BulkWrite(context.Background(), batch)
			if err != nil {
				log.Printf("Failed to insert final batch: %v", err)
			}
		}(operations)
	}

	wg.Wait()
	log.Println("Successfully added products")
}

func InsertingUsersByCopyingFile(db *mongo.Database, fileUri string) {
	file, err := os.Open(fileUri)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	collection := db.Collection("users")
	scanner := bufio.NewScanner(file)
	var operations []mongo.WriteModel
	batchSize := 10000
	var wg sync.WaitGroup
	sem := make(chan struct{}, 100)
	for scanner.Scan() {
		if len(operations) >= batchSize {
			wg.Add(1)
			sem <- struct{}{}
			go func(batch []mongo.WriteModel) {
				defer wg.Done()
				_, err := collection.BulkWrite(context.Background(), batch)
				if err != nil {
					log.Printf("Failed to insert batch: %v", err)
				}
				<-sem
			}(operations)
			operations = nil
		}
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			continue
		}
		userData := strings.Split(line, ";")
		if len(userData) != 2 {
			continue
		}
		userDoc := mongo.NewInsertOneModel().SetDocument(map[string]interface{}{
			"first_name": userData[0],
			"last_name":  userData[1],
		})
		operations = append(operations, userDoc)

	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	wg.Wait()
	log.Println("Successfully added users")
}
