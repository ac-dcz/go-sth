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
	clientOpts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverOpts).SetMaxPoolSize(20)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Println(err)
	}
	defer client.Disconnect(ctx)
	coll := client.Database(dbName).Collection(collName)
	Find(ctx, coll)
}

func Find(ctx context.Context, coll *mongo.Collection) {
	//find all
	var results []bson.D
	if cur, err := coll.Find(ctx, bson.D{}); err != nil {
		log.Println(err)
	} else {
		if err := cur.All(ctx, &results); err != nil {
			log.Println(err)
		} else {
			for _, result := range results {
				if data, err := bson.MarshalExtJSON(result, false, false); err != nil {
					log.Println(err)
				} else {
					log.Printf("%s\n", data)
				}
			}
		}
	}

	//find
	findOpts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 0}})
	if cursor, err := coll.Find(ctx, bson.D{
		{Key: "nick_name", Value: bson.D{{Key: "$regex", Value: "dcz"}}},
	}, findOpts); err != nil {
		log.Println(err)
	} else {
		for cursor.Next(ctx) {
			var user User
			if err := cursor.Decode(&user); err != nil {
				log.Println(err)
			} else {
				log.Println(user)
			}
		}
	}
}
