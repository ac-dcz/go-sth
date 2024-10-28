package main

import (
	"context"
	"log"
	"math/rand"
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
	serverOpts := options.ServerAPI(options.ServerAPIVersion1)
	clientOpts := options.Client().ApplyURI(uri).SetMaxPoolSize(20).SetTimeout(time.Second * 5).SetServerAPIOptions(serverOpts)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Disconnect(ctx)
	coll := client.Database(dbName).Collection(collName)
	Update(coll)
}

func Update(coll *mongo.Collection) {
	ctx := context.TODO()
	//nick_name ___
	var result bson.D
	fuOpts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(bson.M{"_id": 0})
	if err := coll.FindOneAndUpdate(ctx, bson.M{"nick_name": bson.M{"$regex": "^.{3}$"}}, bson.M{"$inc": bson.M{"age": -1}}, fuOpts).Decode(&result); err != nil {
		log.Println(err)
	} else {
		data, _ := bson.MarshalExtJSONIndent(result, false, false, "", "  ")
		log.Println(string(data))
	}

	//updateMany
	if r, err := coll.Find(ctx, bson.D{}); err != nil {
		log.Println(err)
	} else {
		for r.Next(ctx) {
			var obj bson.M
			r.Decode(&obj)
			if r, err := coll.UpdateOne(ctx, bson.M{"_id": obj["_id"]}, bson.M{"$set": bson.M{"desc": bson.A{bson.M{"tel": rand.Int63()}, bson.M{"number": rand.Intn(10)}}}}); err != nil {
				log.Println(err)
			} else {
				log.Printf("%#v \n", r)
			}
		}
	}

}
