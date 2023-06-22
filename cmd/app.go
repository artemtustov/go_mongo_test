package main

import (
	"fmt"
	"go_mongo_test/internal/database"
	"go_mongo_test/internal/transport"
)

func main() {
	fmt.Println("App started")
	database.NewMongoConnection()
	defer database.CloseConnection()
	database.InitCollection()

	transport.NewServer()
}
