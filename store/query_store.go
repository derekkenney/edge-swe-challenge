package store

import (
	"context"
	"fmt"
	"log"

	"github.com/derekkenney/edge-swe-challenge/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewQueryStore(client *mongo.Client) *QueryStore {
	return &QueryStore{
		mongoClient: client,
	}
}

type QueryStore struct {
	mongoClient *mongo.Client
}

func (q *QueryStore) SaveSportEvent(sportsEventMessage *pb.Event) {
	log.Println("SaveSportEvent()")
	collection := q.mongoClient.Database("events").Collection("messages")

	insertResult, err := collection.InsertOne(context.TODO(), sportsEventMessage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
