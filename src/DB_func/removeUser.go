package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func remove(user string, pass string) bool {
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("CAE").Collection("usuarios")
	result, err := collection.DeleteOne(ctx, bson.M{"user": user, "password": pass})
	if err != nil {
		log.Fatal(err)
		return false
	}
	if result.DeletedCount == 0 {
		return false
	}
	return true
}
func main() {
	println(remove("David", "1234"))
}
