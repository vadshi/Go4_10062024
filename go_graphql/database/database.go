package database

import (
	"context"
	"fmt"
	"go_graphql/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func timePtr(t time.Time) *time.Time {
	return &t
}

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	// Get client to work to mongo server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%T\n", client)

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	return &DB {
		client: client,
	}
}

func collectionHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("blog_posts").Collection(collectionName)
}

func(db *DB) GetPost(id string) *model.Post {
	collection := collectionHelper(db, "posts")
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	var post model.Post

	err := collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}
	return &post
}


func(db *DB) GetPosts() []*model.Post {
	collection := collectionHelper(db, "posts")
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	var posts []*model.Post
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &posts); err != nil {
		log.Fatal(err)
	}
	return posts
}

func(db *DB) CreatePost(postInfo *model.NewPost) *model.Post {
	collection := collectionHelper(db, "posts")
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	postInfo.PublishedAt = timePtr(time.Now())
	postInfo.UpdatedAt = timePtr(time.Now())

	result, err := collection.InsertOne(ctx, postInfo)
	if err != nil {
		log.Fatal(err)
	}

	newPost := &model.Post{
		ID: result.InsertedID.(primitive.ObjectID).Hex(),
		Title: postInfo.Title,
		Content: postInfo.Content,
		Author: *postInfo.Author,
		Hero: *postInfo.Hero,
		PublishedAt: *postInfo.PublishedAt,
		UpdatedAt: *postInfo.UpdatedAt,
	}
	return newPost
}

// func(db *DB) UpdatePost(id string, postInfo model.NewPost) *model.Post {
// 	collection := collectionHelper(db, "posts")
// 	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
// 	defer cancel()

// }


// func(db *DB) DeletePost(id string) *model.DeletePostResponse {
// 	collection := collectionHelper(db, "posts")
// 	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
// 	defer cancel()

// }