package ConnectApi

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//clientOption := options.Client().ApplyURI("mongodb://admin:f6XPinsVTx@localhost:27017")
//client, err := mongo.Connect(context.TODO(), clientOption)
//if err != nil {
//log.Fatal(err)
//}
//log.Println("MongoDB connection success")
//collection = client.Database("Books").Collection("Books")

func ConnectDB() *mongo.Client {
	Mongo_URL := "mongodb://admin:f6XPinsVTx@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB")
	return client
}