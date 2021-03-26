package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const mongoUrl = `mongodb://renhao:renhao666@122.9.41.45:27017/admin?authSource=admin`

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))

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
