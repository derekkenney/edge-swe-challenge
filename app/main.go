package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/derekkenney/edge-swe-challenge/business/data/execution"
	"github.com/derekkenney/edge-swe-challenge/business/data/sportsevent"
	pub "github.com/derekkenney/edge-swe-challenge/foundation/pubsub"
	nats "github.com/nats-io/nats.go"
)

const (
	natsServer = "localhost:4222"
	event      = "message-received"
	aggregate  = "message"
	grpcURI    = "localhost:50051"
)

func main() {
	// Subscribe to channels for sport-event, and executions.
	// When message is received, then call the appropriate gRPC function
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	// create an execution object from proto, and send to queue
	executionSaveCommand := execution.ExecutionSaveCommand{
		Symbol:         "AAA",
		Market:         "Philadelphia",
		Price:          5.05,
		Quantity:       10,
		ExecutionEpoch: time.Now().Unix(),
		StateSymbol:    "PA",
	}

	sportSaveCommand := sportsevent.Event{
		Sport:      sportsevent.Sport_BASKETBALL,
		MatchTitle: "NCAA",
		DataEvent:  "Description",
	}

	// Call data.pub. This is a work around since the feed isn't working
	pub.Pub(nc, &executionSaveCommand, &sportSaveCommand)

	// Subscribe to sport_event queue, and store message
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("sport_event", func(m *nats.Msg) {

		// Decode the queue message into event type
		event := sportsevent.Event{}
		// Unmarshal JSON that represents the Event data
		err := json.Unmarshal(m.Data, &event)
		if err != nil {
			log.Print(err)
			return
		}

		log.Print(event)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()
}
