package room

import (
	"log"
	"server/client"
	"server/utils"
)

// Client TODO
type Client = *client.Client

// Clients TODO
type Clients = map[Client]bool

// ClientChannel TODO
type ClientChannel = chan Client

// Room TODO
type Room struct {
	ID string `json:"id"`

	Title string `json:"title"`

	Clients []string `json:"clients"`

	// Registered clients
	clients Clients

	// Inbound messages from clients
	broadcast chan []byte

	// Register requests from clients
	register ClientChannel

	// Unregister requests from clients
	unregister ClientChannel
}

// New TODO
func New(ID string, Title string) *Room {
	return &Room{
		ID:      ID,
		Title:   Title,
		Clients: []string{},

		clients:    Clients{},
		broadcast:  make(chan []byte),
		register:   make(ClientChannel),
		unregister: make(ClientChannel),
	}
}

// Start TODO
func (room *Room) Start(onChange func(room *Room)) {
	for {
		select {

		case client := <-room.register:
			room.addClient(client)

			log.Printf("client %v join room %v ", client.ID, room.ID)

			onChange(room)

		case client := <-room.unregister:
			if !room.clients[client] {
				return
			}

			room.removeClient(client)

			log.Printf("client %v leave room %v ", client.ID, room.ID)

			onChange(room)

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

// Join TODO
func (room *Room) Join() ClientChannel {
	return room.register
}

// Leave TODO
func (room *Room) Leave() ClientChannel {
	return room.unregister
}

// Has TODO
func (room *Room) Has(target Client) bool {
	return room.clients[target]
}

func (room *Room) addClient(client Client) {
	room.clients[client] = true
	room.Clients = append(room.Clients, client.ID)
	client.RoomID = room.ID
}

func (room *Room) removeClient(client Client) {
	delete(room.clients, client)
	room.Clients = utils.Remove(room.Clients, client.ID)
	client.RoomID = ""
}
