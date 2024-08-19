package tasks

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateReports(db *mongo.Database) {
	fmt.Println("Total Discount Received by Each User:")
	discountPipeline := mongo.Pipeline{
		{{"$lookup", bson.D{
			{"from", "orders"},
			{"localField", "_id"},
			{"foreignField", "user_id"},
			{"as", "user_orders"},
		}}},
		{{"$unwind", "$user_orders"}},
		{{"$group", bson.D{
			{"_id", "$_id"},
			{"total_discount", bson.D{{"$sum", "$user_orders.discount"}}},
		}}},
	}
	cursor, err := db.Collection("users").Aggregate(context.Background(), discountPipeline)
	if err != nil {
		log.Fatalf("Failed to aggregate discounts: %v", err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			log.Fatalf("Failed to decode discount result: %v", err)
		}
		fmt.Printf("UserID: %v, TotalDiscount: %v\n", result["_id"], result["total_discount"])
	}

	fmt.Println("Total Amount of Each Order:")
	amountPipeline := mongo.Pipeline{
		{{"$group", bson.D{
			{"_id", "$order_id"},
			{"total_amount", bson.D{{"$sum", "$amount"}}},
		}}},
	}
	cursor, err = db.Collection("orders").Aggregate(context.Background(), amountPipeline)
	if err != nil {
		log.Fatalf("Failed to aggregate order amounts: %v", err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			log.Fatalf("Failed to decode order amount result: %v", err)
		}
		fmt.Printf("OrderID: %v, TotalAmount: %v\n", result["_id"], result["total_amount"])
	}

	fmt.Println("Goods Received by Each User:")
	goodsPipeline := mongo.Pipeline{
		{{"$lookup", bson.D{
			{"from", "order_products"},
			{"localField", "order_id"},
			{"foreignField", "order_id"},
			{"as", "products"},
		}}},
		{{"$unwind", "$products"}},
		{{"$group", bson.D{
			{"_id", "$user_id"},
			{"goods", bson.D{{"$addToSet", "$products.product_name"}}},
		}}},
	}
	cursor, err = db.Collection("orders").Aggregate(context.Background(), goodsPipeline)
	if err != nil {
		log.Fatalf("Failed to aggregate goods: %v", err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			log.Fatalf("Failed to decode goods result: %v", err)
		}
		fmt.Printf("UserID: %v, Goods: %v\n", result["_id"], result["goods"])
	}

	fmt.Println("Branches From Which Goods Were Received:")
	branchPipeline := mongo.Pipeline{
		{{"$lookup", bson.D{
			{"from", "order_products"},
			{"localField", "order_id"},
			{"foreignField", "order_id"},
			{"as", "products"},
		}}},
		{{"$unwind", "$products"}},
		{{"$lookup", bson.D{
			{"from", "branches"},
			{"localField", "products.branch_id"},
			{"foreignField", "_id"},
			{"as", "branch_info"},
		}}},
		{{"$unwind", "$branch_info"}},
		{{"$group", bson.D{
			{"_id", "$user_id"},
			{"branches", bson.D{{"$addToSet", "$branch_info.branch_name"}}},
		}}},
	}
	cursor, err = db.Collection("orders").Aggregate(context.Background(), branchPipeline)
	if err != nil {
		log.Fatalf("Failed to aggregate branches: %v", err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			log.Fatalf("Failed to decode branches result: %v", err)
		}
		fmt.Printf("UserID: %v, Branches: %v\n", result["_id"], result["branches"])
	}
}
