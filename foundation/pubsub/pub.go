package pub

import (
	"encoding/json"
	"log"

	"github.com/derekkenney/edge-swe-challenge/business/data/execution"
	"github.com/derekkenney/edge-swe-challenge/business/data/sportsevent"
	"github.com/nats-io/nats.go"
)

func Pub(nc *nats.Conn, execution *execution.ExecutionSaveCommand, event *sportsevent.Event) {
	go publishExecution(nc, execution)
	go publishSportEvent(nc, event)
	log.Println("Published message to ", nats.DefaultURL)
}

func publishExecution(nc *nats.Conn, execution *execution.ExecutionSaveCommand) {
	subject := "execution"
	data, err := json.Marshal(&execution)

	if err != nil {
		log.Fatal(err)
	}
	// Publish message on subject
	nc.Publish(subject, data)
	log.Println("Published message on subject " + subject)
}

func publishSportEvent(nc *nats.Conn, event *sportsevent.Event) {
	subject := "sport_event"
	data, err := json.Marshal(&event)

	if err != nil {
		log.Fatal(err)
	}
	// Publish message on subject
	nc.Publish(subject, data)
	log.Println("Published message on subject " + subject)
}
