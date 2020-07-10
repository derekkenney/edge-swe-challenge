package main

import (
	"encoding/json"
	"log"

	"github.com/derekkenney/edge-swe-challenge/pb"
	"github.com/derekkenney/edge-swe-challenge/store"
	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/nats.go"
)

func main() {
	mongoClient := store.NewMongoClient()

	// Connect to NATS server
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()

	// create a message from sportevent type
	sportEventMessage := pb.Event{
		Sport:      pb.Sport_BASKETBALL,
		MatchTitle: "March Madness",
		DataEvent:  "March Madness event description",
	}

	eventData, err := proto.Marshal(&sportEventMessage)
	subject := "sport_event"
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	// Publish message on subject
	natsConnection.Publish(subject, eventData)
	log.Println("Published message on subject " + subject)

	log.Print("saveSportEventMessage()")
	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		sportEventMessage := pb.Event{}
		err := json.Unmarshal(msg.Data, &sportEventMessage)
		if err != nil {
			log.Print(err)
			return
		}
	})

	log.Println(sportEventMessage)

	// Handle the message
	store := store.NewQueryStore(mongoClient)
	store.SaveSportEvent(&sportEventMessage)
}
