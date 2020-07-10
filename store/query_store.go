package store

import (
	"context"
	"fmt"
	"log"

	"github.com/derekkenney/edge-swe-challenge/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueryStore struct{}

func (q *QueryStore) SaveSportEvent(sportsEventMessage *pb.Event) {
	log.Println("SaveSportEvent()")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("events").Collection("messages")

	insertResult, err := collection.InsertOne(context.TODO(), sportsEventMessage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
