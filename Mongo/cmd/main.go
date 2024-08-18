package main

import (
	"awesomeProject/Mongo/mongo"
	"context"
	"fmt"
	"log"
)

func main() {
	uri := "mongodb://localhost:27017"
	dbName := "test"
	client, err := mongo.ConnectToMongoDB(uri, dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Printf("database name %v \n", mongo.Database.Name())
}
