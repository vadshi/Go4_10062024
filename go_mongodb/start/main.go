package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Insert many documents
	examples := []interface{}{
		bson.D{
			{Key: "someString", Value: "Second String"},
			{Key: "someInteger", Value: 121},
			{Key: "someStringSlice", Value: []string{"Example11", "Example12", "Example13"}},
		},
		bson.D{
			{Key: "someString", Value: "Another example String"},
			{Key: "someInteger", Value: 19},
			{Key: "someStringSlice", Value: []string{"Example21", "Example22"}},
		},
	}

	rs, err := exampleCollection.InsertMany(ctx, examples)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs.InsertedIDs)

	// find document by ObjectID
	c := exampleCollection.FindOne(ctx, bson.M{"_id": r.InsertedID})

	var exampleResult bson.M
	err = c.Decode(&exampleResult)  
	if err != nil {
		log.Fatal(err)
	}
	// print document data
	fmt.Printf("\nItem with ID: %v containing the following:\n", exampleResult["_id"])
	fmt.Println("someString", exampleResult["someString"])
	fmt.Println("someInteger", exampleResult["someInteger"])
	fmt.Println("someStringSlice", exampleResult["someStringSlice"])

	// find document by value of ObjectID
	objectId, err := primitive.ObjectIDFromHex("6666ccb73649a6cc6fb508f3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(objectId)

	new_c := exampleCollection.FindOne(ctx, bson.M{"_id": bson.M{"$eq": objectId}})

	var exampleRes bson.M
	fmt.Println("\nresult type", reflect.TypeOf(exampleRes))
	fmt.Println("result before", exampleRes)

	err = new_c.Decode(&exampleRes)  
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After:")
	// print document data
	fmt.Printf("\nItem with ID: %v containing the following:\n", exampleRes["_id"])
	fmt.Println("someString", exampleRes["someString"])
	fmt.Println("someInteger", exampleRes["someInteger"])
	fmt.Println("someStringSlice", exampleRes["someStringSlice"])


}