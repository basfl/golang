package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://{user}:{pass}@cluster0-mmvmw.mongodb.net/golang?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal("connection error occured")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(ctx)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("connection error occured")
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("connection error occured")
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal("connection error occured")
	}
	fmt.Println(databases)
	testDatabase := client.Database("test")
	podcastsCollection := testDatabase.Collection("podcats")
	episodsCollection := testDatabase.Collection("espisods")
	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{"title", "title1"},
		{"author", "basfl"},
		{"tags", bson.A{"dev", "int", "pal"}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastResult.InsertedID)
	episodResult, err := episodsCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"titile", "ep1"},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"titile", "ep2"},
			{"description", "description 2"},
			{"duration", 25},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodResult.InsertedIDs)
	//reading all the values in collection
	cursor, err := episodsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// var episods []bson.M
	// if err = cursor.All(ctx, &episods); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(episods)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episod bson.M
		if err = cursor.Decode(&episod); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episod)
	}
	fmt.Println("\n ******* read just one ***** \n")

	var podcast bson.M
	filter := bson.M{"title": "title1"}
	if err = podcastsCollection.FindOne(ctx, filter).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcast)

}
