package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri      = "mongodb://127.0.0.1:27017/?timeoutMS=5000"
	dbName   = "test"
	collName = "user_info"
)

type User struct {
	NickName string `bson:"nick_name"`
	Age      int    `bson:"age"`
	Address  struct {
		Country string `bson:"country"`
		City    string `bson:"city"`
	}
	Tags []any `bson:"tags"`
}

func main() {
	serverOpts := options.ServerAPI(options.ServerAPIVersion1)
	clientOpts := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverOpts).
		SetMaxPoolSize(20)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalln(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalln(err)
	}
	defer client.Disconnect(ctx)
	coll := client.Database("test").Collection("user_info")
	if n, err := coll.CountDocuments(ctx, bson.D{}); err != nil {
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Documents: %d\n", n)
	}
	Insert(ctx, coll)
}

func Insert(ctx context.Context, coll *mongo.Collection) {
	//insertOne with bson.D
	if result, err := coll.InsertOne(ctx, bson.D{
		{Key: "nick_name", Value: "dcz-1"},
		{Key: "age", Value: 24},
		{Key: "address", Value: bson.D{{Key: "country", Value: "China"}, {Key: "city", Value: "Wuhan"}}},
		{Key: "tags", Value: bson.A{"red", "male"}},
	}); err != nil {
		log.Println(err)
	} else {
		log.Println(result.InsertedID)
	}

	//insertOne with struct
	sg := User{
		NickName: "sg-1",
		Age:      23,
		Address: struct {
			Country string "bson:\"country\""
			City    string "bson:\"city\""
		}{
			Country: "China",
			City:    "Shanghai",
		},
		Tags: []any{
			"pink",
			"female",
		},
	}

	if result, err := coll.InsertOne(ctx, sg); err != nil {
		log.Println(err)
	} else {
		log.Println(result.InsertedID)
	}

	//insertMany
	inmOpts := options.InsertMany().SetOrdered(true)
	if result, err := coll.InsertMany(ctx, bson.A{
		bson.D{
			{Key: "nick_name", Value: "lhd-1"},
			{Key: "age", Value: 23},
			{Key: "address", Value: bson.D{{Key: "country", Value: "China"}, {Key: "city", Value: "Shengzheng"}}},
			{Key: "tags", Value: bson.A{"yellow", "male"}},
		},
		bson.D{
			{Key: "nick_name", Value: "pyp-1"},
			{Key: "age", Value: 22},
			{Key: "address", Value: bson.D{{Key: "country", Value: "China"}, {Key: "city", Value: "Wuhan"}}},
			{Key: "tags", Value: bson.A{"black", "male"}},
		},
	}, inmOpts); err != nil {
		log.Println(err)
	} else {
		for _, id := range result.InsertedIDs {
			log.Printf("%v %T\n", id, id)
		}
	}

}
