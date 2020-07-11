package main

import (
	"log"
	"time"

	"github.com/derekkenney/edge-swe-challenge/pb"

	//"github.com/derekkenney/edge-swe-challenge/store"

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

	sportChanSend := make(chan *pb.Event)
	ec.BindSendChan("sport_event", sportChanSend)

	i := 0
	for {
		sportEventMessage := pb.Event{
			Sport:        pb.Sport_BASKETBALL,
			MatchTitle:   "March Madness",
			DataEvent:    "March Madness event description",
			CreationDate: time.Now().Unix(),
		}

		// Just send to the channel! :)
		log.Printf("Sending request %d", i)
		sportChanSend <- &sportEventMessage

		// Pause and increment counter
		time.Sleep(time.Second * 1)
		i = i + 1

	}
}
