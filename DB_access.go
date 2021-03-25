package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = "mongodb://localhost:27017"

func addOferta(id_user string, producto string, vendedor string, total int, cuota int, periodo int, cae float64) bool {
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
	collection := client.Database("CAE").Collection("ofertas")
	res, err := collection.InsertOne(ctx, bson.D{{"id_user", id_user}, {"producto", producto}, {"vendedor", vendedor}, {"total", total}, {"cuota", cuota}, {"periodo", periodo}, {"cae", cae}})
	if err != nil {
		return false
	}
	println(res)
	return true
}

func getOfertas(user string) []bson.M {
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
	collection := client.Database("CAE").Collection("ofertas")
	filterCursor, err := collection.Find(ctx, bson.M{"id_user": user})
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return (episodesFiltered)
}

func removeOferta(id string) bool {
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

	collection := client.Database("CAE").Collection("ofertas")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	} else {
		res, err := collection.DeleteOne(ctx, bson.M{"_id": idPrimitive})

		if err != nil {
			log.Fatal("DeleteOne() ERROR:", err)
		} else {
			// Check if the response is 'nil'
			if res.DeletedCount == 0 {
				return false
			} else {
				return true
			}
		}
	}
	return true
}

func addUser(user string, password string) bool {
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
	res, err := collection.InsertOne(ctx, bson.D{{"user", user}, {"password", password}})
	if err != nil {
		return false
	}
	println(res)
	return true
}

func existUser(username string) bool {
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
	filterCursor, err := collection.Find(ctx, bson.M{"user": username})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	if episodesFiltered != nil {
		return true
	}
	if err != nil {
		log.Fatal(err)
	}
	return false
}

func loginCheck(user string, pass string) bool {
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
	filterCursor, err := collection.Find(ctx, bson.M{"user": user, "password": pass})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	if episodesFiltered != nil {
		return true
	}
	if err != nil {
		log.Fatal(err)
	}
	return false
}
func ofCheck(user string, pass string) bool {
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
	filterCursor, err := collection.Find(ctx, bson.M{"user": user, "password": pass})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	if episodesFiltered != nil {
		return true
	}
	if err != nil {
		log.Fatal(err)
	}
	return false
}
