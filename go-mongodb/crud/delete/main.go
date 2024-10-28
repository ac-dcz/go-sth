package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri      = "mongodb://127.0.0.1:27017"
	dbName   = "test"
	collName = "user_info"
)

func main() {
	ctx := context.Background()
	serverOptS := options.ServerAPI(options.ServerAPIVersion1)
	clientOpts := options.Client().ApplyURI(uri).SetTimeout(time.Second * 5).SetServerAPIOptions(serverOptS)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println(err)
	}
	defer client.Disconnect(ctx)

	Delete(client.Database(dbName).Collection(collName))
}

func Delete(coll *mongo.Collection) {
	ctx := context.Background()
	if r, err := coll.DeleteOne(ctx, bson.M{"nick_name": "pyp-1"}); err != nil {
		log.Println(err)
	} else {
		log.Printf("%#v \n", *r)
	}
}
