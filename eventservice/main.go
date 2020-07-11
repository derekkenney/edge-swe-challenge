package main

import (
	"log"

	"github.com/derekkenney/edge-swe-challenge/pb"
	"github.com/derekkenney/edge-swe-challenge/store"

	//"github.com/derekkenney/edge-swe-challenge/store"

	nats "github.com/nats-io/nats.go"
)

func main() {
	mongoClient := store.NewMongoClient()
	store := store.NewQueryStore(mongoClient)
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
		sportReq := <-sportChanRecv
		executionReq := <-executionChanRecv

		// Will execute each function call concurrently in own thread
		go store.SaveSportEvent(sportReq)
		go store.SaveExecutionEvent(executionReq)
	}
}
