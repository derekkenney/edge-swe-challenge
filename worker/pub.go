package main

import (
	"log"
	"time"

	"github.com/derekkenney/edge-swe-challenge/pb"
	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	log.Println("Connected to NATS and ready to send messages")

	// For sports events
	sportChanSend := make(chan *pb.Event)
	ec.BindSendChan("sport_event", sportChanSend)
	// For executions
	executionChanSend := make(chan *pb.Execution)
	ec.BindSendChan("execution", executionChanSend)

	i := 0
	for {
		sportEventMessage := pb.Event{
			Sport:        pb.Sport_BASKETBALL,
			MatchTitle:   "March Madness",
			DataEvent:    "March Madness event description",
			CreationDate: time.Now().Unix(),
		}

		// Just send to the channel! :)
		log.Printf("Sending Sport Event request %d", i)
		sportChanSend <- &sportEventMessage

		executionEventMessage := pb.Execution{
			Symbol:         "AAA",
			Market:         "Boston",
			Price:          10.05,
			Quantity:       10,
			ExecutionEpoch: time.Now().Unix(),
			StateSymbol:    "MA",
		}

		// Just send to the channel! :)
		log.Printf("Sending Execution request %d", i)
		executionChanSend <- &executionEventMessage

		// Pause and increment counter
		time.Sleep(time.Second * 1)
		i = i + 1
	}
}
