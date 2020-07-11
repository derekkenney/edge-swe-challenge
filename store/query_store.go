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
	collection := q.mongoClient.Database("events").Collection("sports")

	insertResult, err := collection.InsertOne(context.TODO(), sportsEventMessage)
	if err != nil {
		log.Printf("SaveSportEvent() error. Couldn't insert record to DB %v\n", err)
		return
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func (q *QueryStore) SaveExecutionEvent(executionEventMessage *pb.Execution) {
	log.Println("SaveExecutionEvent()")
	collection := q.mongoClient.Database("events").Collection("executions")

	insertResult, err := collection.InsertOne(context.TODO(), executionEventMessage)
	if err != nil {
		log.Printf("SaveExecutionEvent() error. Couldn't insert record to DB %v\n", err)
		return
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
