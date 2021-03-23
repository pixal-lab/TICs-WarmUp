package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Replace the uri string with your MongoDB deployment's connection string.
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

	collection := client.Database("escuela").Collection("alumnos")
	//Ingresar a base de datos
	/*res, err := collection.InsertOne(ctx, bson.D{{"nombre", "Pablo"}, {"edad", 22},{"clave","123"},{"grado","egresado"}})
	fmt.Println(res.InsertedID)*/
	// Ping the primary
	//Imprimir a todos
	/*cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episode)
	}*/
	//buscar 1 usuario
	/*filterCursor, err := collection.Find(ctx, bson.M{"clave": "123"})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)*/
	//Borrar (funar)
	/*result, err := collection.DeleteOne(ctx, bson.M{"nombre": "Jhon"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)*/
	/*if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	//filterCursor, err := collection.FindId(bson.M{"_id": bson.ObjectIdHex("6057917d3658c1a684d12962")})
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)*/
	// Declare a primitive ObjectID from a hexadecimal string
	idPrimitive, err := primitive.ObjectIDFromHex("605790853658c1a684d12961")
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	} else {
		res, err := collection.DeleteOne(ctx, bson.M{"_id": idPrimitive})
		fmt.Println("DeleteOne Result TYPE:", reflect.TypeOf(res))

		if err != nil {
			log.Fatal("DeleteOne() ERROR:", err)
		} else {
			// Check if the response is 'nil'
			if res.DeletedCount == 0 {
				fmt.Println("DeleteOne() document not found:", res)
			} else {
				// Print the results of the DeleteOne() method
				fmt.Println("DeleteOne Result:", res)

				// *mongo.DeleteResult object returned by API call
				fmt.Println("DeleteOne TYPE:", reflect.TypeOf(res))
			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected and pinged.")
}
