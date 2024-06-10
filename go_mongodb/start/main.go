package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	// Get client to work to mongo server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	
	// Close connection
	defer client.Disconnect(ctx)

	fmt.Printf("%T\n", client)

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	// Get all database names
	dbNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbNames)

	// Create new database and collection
	testDB := client.Database("test")

	fmt.Printf("%T\n", testDB)

	exampleCollection := testDB.Collection("example")

	// defer exampleCollection.Drop(ctx)

	fmt.Printf("%T\n", exampleCollection)

	// Insert new document in db
	example := bson.D{
		{Key: "someString", Value: "Example String"},
		{Key: "someInteger", Value: 12},
		{Key: "someStringSlice", Value: []string{"Example1", "Example2", "Example3"}},
	}

	r, err := exampleCollection.InsertOne(ctx, example)
	if err != nil {
		log.Fatal(err)
	}
	// print "_id" of new document
	fmt.Println(r.InsertedID)

}