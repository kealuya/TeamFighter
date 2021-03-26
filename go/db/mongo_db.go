package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"team_fighter_go/conf"
	"time"
)

var mongoClient *mongo.Client
var once sync.Once
var ctx = context.Background()

func ObtainMongoCollection(collection string) (*mongo.Collection, context.Context) {

	once.Do(func() {
		url := conf.GetConfigWithKey("db.mongoUrl")
		ctx := context.Background()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
		if err != nil {
			log.Panicln(err)
		}
		mongoClient = client
	})
	return mongoClient.Database("team_fighter").Collection(collection), ctx
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongoUrl"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("admin").Collection("test")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(result)
		resultMap := result.Map()
		log.Println(fmt.Sprintf("%T", resultMap["_id"]))
		log.Println(resultMap["user"])
		log.Println(resultMap["age"])
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
