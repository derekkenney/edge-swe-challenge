package store

import (
	"context"
	"fmt"
	"log"

	"github.com/derekkenney/edge-swe-challenge/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Types should have a single purpose. It allows our programs to change and grow as data transformation changes
// Saving events is common behavior. Types are grouped according to that commong behavior. This is composition of
// common types.

type ExecutionQueryStore struct {
	Client    *mongo.Client
	Execution *pb.Execution
}

type SportQueryStore struct {
	Client *mongo.Client
	Event  *pb.Event
}

type ChatQueryStore struct {
	Client  *mongo.Client
	Message *pb.ChatMessage
}

// Save sport events to data store
func (q *SportQueryStore) Save() {
	log.Println("SaveSportEvent()")
	collection := q.Client.Database("events").Collection("sports")

	insertResult, err := collection.InsertOne(context.TODO(), q.Event)
	if err != nil {
		log.Printf("SaveSportEvent() error. Couldn't insert record to DB %v\n", err)
		return
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// Save() Persists execution message to DB
func (q *ExecutionQueryStore) Save() {
	log.Println("SaveExecutionEvent()")
	collection := q.Client.Database("events").Collection("executions")

	insertResult, err := collection.InsertOne(context.TODO(), q.Execution)
	if err != nil {
		log.Printf("SaveExecutionEvent() error. Couldn't insert record to DB %v\n", err)
		return
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// Save() Persists chat messages to DB
func (q *ChatQueryStore) Save() {
	log.Println("SaveChatMessageEvent()")
	collection := q.Client.Database("events").Collection("messages")

	insertResult, err := collection.InsertOne(context.TODO(), q.Message)
	if err != nil {
		log.Printf("SaveChatMessageEvent() error. Couldn't insert record to DB %v\n", err)
		return
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
