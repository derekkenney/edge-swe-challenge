package main

import (
	"log"

	"github.com/derekkenney/edge-swe-challenge/pb"
	"github.com/derekkenney/edge-swe-challenge/store"

	//"github.com/derekkenney/edge-swe-challenge/store"

	nats "github.com/nats-io/nats.go"
)

func main() {
	var test string
	mongoClient := store.NewMongoClient()
	sportStore := store.SportQueryStore{
		Client: mongoClient,
	}
	executionStore := store.ExecutionQueryStore{
		Client: mongoClient,
	}

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
	executionChanRecv := make(chan *pb.Execution)
	ec.BindRecvChan("execution", executionChanRecv)

	for {
		// Wait for incoming messages
		sportEvent := <-sportChanRecv
		executionEvent := <-executionChanRecv

		// Will execute each function call concurrently in own thread
		sportStore.Event = sportEvent
		executionStore.Execution = executionEvent

		go sportStore.Save()
		go executionStore.Save()
	}
}
