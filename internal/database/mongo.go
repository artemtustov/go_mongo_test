package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_mongo_test/internal/domain"
	"log"
	"time"
)

var Client *mongo.Client
var Collection *mongo.Collection

func NewMongoConnection() {
	cont, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(cont, clientOptions)
	if err != nil {
		log.Fatal("Can't connect to MongoDB!")
	}

	err = client.Ping(cont, nil)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	fmt.Println("Connected to MongoDB!")
}

func InitCollection() {
	Collection = Client.Database("inventory").Collection("cars")
}

func CloseConnection() {
	cont, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err := Client.Disconnect(cont)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func AddCar(c domain.CarInternal) error {
	insertResult, err := Collection.InsertOne(context.TODO(), c)
	if err != nil {
		return err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

func FindCars() ([]domain.CarInternal, error) {
	var cars []domain.CarInternal

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem domain.CarInternal
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		cars = append(cars, elem)
	}

	fmt.Println("Found multiple documents ")
	return cars, nil
}
