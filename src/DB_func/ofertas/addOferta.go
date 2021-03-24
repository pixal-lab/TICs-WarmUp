package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func addOferta(id_user string, producto string, vendedor string, total int, cuota int, periodo int, cae float64) bool {
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
	collection := client.Database("CAE").Collection("ofertas")
	res, err := collection.InsertOne(ctx, bson.D{{"id_user", id_user}, {"producto", producto}, {"vendedor", vendedor}, {"total", total}, {"cuota", cuota}, {"periodo", periodo}, {"cae", cae}})
	if err != nil {
		return false
	}
	println(res)
	return true
}

/*func main() {
	println(addOferta("Catalina", "Spalding NBA Marble", "Nico Bolton", 25000, 100, 3, 3.23))
}*/
