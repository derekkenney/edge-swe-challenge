package main

import (
	"log"
	"net/http"
	"time"

	"github.com/derekkenney/edge-swe-challenge/store"

	"github.com/derekkenney/edge-swe-challenge/pb"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *pb.ChatMessage)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fs := http.FileServer(http.Dir("../public"))
	store := store.NewQueryStore(store.NewMongoClient())

	http.Handle("/", fs)
	http.Handle("/ws", handleConnections(store))

	go handleMessages(store)

	log.Println("http server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleConnections(store *store.QueryStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer ws.Close()

		clients[ws] = true

		for {
			var msg *pb.ChatMessage
			// Read in a new message as JSON and map it to a ChatMessage object
			err := ws.ReadJSON(&msg)
			if err != nil {
				log.Printf("error: %v", err)
				delete(clients, ws)
				break
			}
			broadcast <- msg

			//TODO: Persist reading a message to event store
			msg.CreationDate = time.Now().Unix()
			msg.EventType = "read_message"
			store.SaveChatMessageEvent(msg)
		}
	})
}

func handleMessages(store *store.QueryStore) {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}

			//TODO: Persist writing a message to event store
			msg.CreationDate = time.Now().Unix()
			msg.EventType = "write_message"
			store.SaveChatMessageEvent(msg)
		}
	}
}
