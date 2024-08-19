package test

import (
	"awesomeProject/Mongo/tasks"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

func setupTestDb() (*mongo.Database, func()) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	db := client.Database("test")
	cleanup := func() {
		_ = client.Disconnect(context.Background())
	}
	return db, cleanup
}

func BenchmarkInsertingBranches(b *testing.B) {
	//db, f := setupTestDb()
	//defer f()
	//for i := 0; i < b.N; i++ {
	//	tasks.InsertingBranchesByIteratingFile(db, "/home/elon/GolandProjects/database_task_go/Mongo/example_data/5k_users.txt")
	//	b.StopTimer()
	//	break
	//}
}

/*
elon@musk:~/GolandProjects/database_task_go/Mongo/test$ go test -bench .
successfully added
goos: linux
goarch: amd64
pkg: awesomeProject/Mongo/test
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkInsertingBranches-16           successfully added
successfully added
successfully added
successfully added
successfully added
successfully added
successfully added
1000000000               0.04684 ns/op
PASS
ok      awesomeProject/Mongo/test       0.438s
*/

func BenchmarkInsertingProducts(b *testing.B) {
	//db, f := setupTestDb()
	//defer f()
	//b.ResetTimer()
	//tasks.InsertingProductByCopyingFile(db, "/home/elon/GolandProjects/database_task_go/Mongo/example_data/1mln_products.txt")
	//b.StopTimer()
}

/*
InsertingBranches-16           1000000000               0.0000002 ns/op
2024/08/18 18:16:11 Successfully added products
BenchmarkInsertingProducts-16                  1        1771716813 ns/op
PASS
ok      awesomeProject/Mongo/test       1.860s
elon@musk:~/GolandProjects/database_task_go/Mongo/test$
*/

//NOT PASSED MAX=1SECOND

func BenchmarkInsertUsers(b *testing.B) {
	db, f := setupTestDb()
	defer f()
	b.ResetTimer()
	tasks.InsertingUsersByCopyingFile(db, "/home/elon/GolandProjects/database_task_go/Mongo/example_data/5k_users.txt")
	b.StopTimer()
}

/*
          0.0004882 ns/op
PASS
ok      awesomeProject/Mongo/test       0.042s
elon@musk:~/GolandProjects/database_task_go/Mongo/test$
*/

//PASSED
