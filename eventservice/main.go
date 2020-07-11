package main

import (
	"log"

	"github.com/derekkenney/edge-swe-challenge/pb"
	"github.com/derekkenney/edge-swe-challenge/store"

	//"github.com/derekkenney/edge-swe-challenge/store"

	nats "github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	mongoClient := store.NewMongoClient()

	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	log.Println("Connected to NATS and ready to receive messages")

	sportChanRecv := make(chan *pb.Event)
	ec.BindRecvChan("sport_event", sportChanRecv)

	//sub(natsConnection)
	//saveMessage(mongoClient, response)
	// Wait for incoming messages

	for {
		// Wait for incoming messages
		req := <-sportChanRecv
		log.Printf("Received request: %v", req)

		// Will execute each function call concurrently in own thread
		go saveMessage(mongoClient, req)
	}
}
func saveMessage(mongoClient *mongo.Client, event *pb.Event) {
	log.Println("saveMessage()")

	store := store.NewQueryStore(mongoClient)
	store.SaveSportEvent(event)
}
