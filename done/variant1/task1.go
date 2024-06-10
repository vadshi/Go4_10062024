package main

import (
	"context"
	"fmt"
	"log"
	//"reflect"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// TODO. Реализовать CRUD для коллекции "Книги"

type Book struct {
	Title     string 	`bson:"title, omitempty"`
	Author    string 	`bson:"author, omitempty"`
	ISBN      string 	`bson:"isbn, omitempty"`
	Publisher string 	`bson:"publisher, omitempty"`
	Year      int		`bson:"year, omitempty"`
	Copies    int		`bson:"copies, omitempty"`
}

type BookObject struct {
	ID   string `bson:"_id"`
	Book1 Book  `bson:"book1"`
}

func main() {
	fmt.Println("DB==========================")
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", client)
	defer client.Disconnect(ctx)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	// Имена БД
	dbNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbNames)

	// Новая БД и коллекция
	fmt.Println("Collection=================")
	testdb := client.Database("books")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", testdb)
	testCollection := testdb.Collection("booksColl")
	if err != nil {
		log.Fatal(err)
	}
	// defer testCollection.Drop(ctx)

	// Вставка многого
	fmt.Println("Create====================")
	books := []interface{}{
		bson.D{
			{Key: "book1", Value: Book{Title: "t1", Author: "a1", ISBN: "i1", Publisher: "p1", Year: 1, Copies: 1}},
		},
		bson.D{
			{Key: "book2", Value: Book{Title: "t2", Author: "a2", ISBN: "i2", Publisher: "p2", Year: 2, Copies: 2}},
		},
	}
	resC, err := testCollection.InsertMany(ctx, books)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resC.InsertedIDs)

	// Поиск по ID
	fmt.Println("Read========================")
	InsertedID := resC.InsertedIDs[0]
	c := testCollection.FindOne(ctx, bson.M{"_id": InsertedID})
	// var resF bson.M
	var resF BookObject
	err2 := c.Decode(&resF)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Item with ID: %v\n", resF.ID)
	fmt.Printf("Item with Author: %v\n", resF.Book1.Author)
	fmt.Printf("Item with Title: %v\n", resF.Book1.Title)

	// Обновление
	fmt.Println("Update======================")
	//c[0].["Title"]="t11"
	rUpd, err := testCollection.UpdateOne(ctx,
		bson.M{"_id": InsertedID},
		bson.D{
			// {Key: "$set", Value: bson.M{"book1": Book{Title: "t11"}}},  // change full document 
			{Key: "$set", Value: bson.M{"book1.title": "t11"}},            // change only one field in "book1"
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ModifiedCount: ", rUpd.ModifiedCount)
	updated_c := testCollection.FindOne(ctx, bson.M{"_id": InsertedID})
	var resU bson.M
	err = updated_c.Decode(&resU)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Item with ID: %v containing:\n", resU["_id"])
	fmt.Println("book1", resU["book1"])
	fmt.Println("Finish")

}
