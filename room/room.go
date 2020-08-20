package room

import (
	"log"
	"server/model"
)

// Room TODO
type Room struct {
	// Registered clients
	clients model.Clients

	// Inbound messages from clients
	broadcast chan []byte

	// Register requests from clients
	register chan model.Client

	// Unregister requests from clients
	unregister chan model.Client
}

// New TODO
func New() model.Room {
	return &Room{
		clients:    model.Clients{},
		broadcast:  make(chan []byte),
		register:   make(chan model.Client),
		unregister: make(chan model.Client),
	}
}

// Start TODO
func (room *Room) Start() {
	for {
		select {

		case client := <-room.register:
			room.clients[client] = true

			log.Printf("number of person: %d\n", len(room.clients))

		case client := <-room.unregister:
			if !room.clients[client] {
				return
			}

			delete(room.clients, client)
			client.Close()

			log.Printf("number of person: %d\n", len(room.clients))

		case msg := <-room.broadcast:

			for client := range room.clients {
				client.Send(msg)
			}
		}
	}
}

// Broadcast TODO
func (room *Room) Broadcast(msg []byte) {
	room.broadcast <- msg
}

// Add TODO
func (room *Room) Add(client model.Client) {
	log.Println("client connected...")

	client.Start(room)

	room.register <- client
}

// Delete TODO
func (room *Room) Delete(client model.Client) {
	room.unregister <- client
}
