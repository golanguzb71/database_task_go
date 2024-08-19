package tasks

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CalculateBranchFinancials(db *mongo.Database) {
	pipeline := mongo.Pipeline{
		{{"$lookup", bson.D{
			{"from", "order_products"},
			{"localField", "order_id"},
			{"foreignField", "order_id"},
			{"as", "order_products"},
		}}},
		{{"$unwind", "$order_products"}},
		{{"$lookup", bson.D{
			{"from", "branches"},
			{"localField", "order_products.branch_id"},
			{"foreignField", "_id"},
			{"as", "branch_info"},
		}}},
		{{"$unwind", "$branch_info"}},
		{{"$group", bson.D{
			{"_id", "$branch_info.branch_name"},
			{"total_income", bson.D{{"$sum", bson.D{
				{"$subtract", bson.A{
					bson.D{{"$sum", "$order_products.outcome_price"}},
					bson.D{{"$sum", "$order_products.discount_price"}},
					bson.D{{"$sum", "$order_products.income_price"}},
				}},
			}}}},
			{"total_expenses", bson.D{{"$sum", "$order_products.discount_price"}}},
		}}},
	}

	cursor, err := db.Collection("orders").Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatalf("Failed to aggregate branch financials: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatalf("Failed to decode result: %v", err)
		}
		fmt.Printf("Branch: %v, Total Income: %v, Total Expenses: %v\n",
			result["_id"], result["total_income"], result["total_expenses"])
	}
}
