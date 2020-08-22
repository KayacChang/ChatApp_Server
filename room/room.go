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

// ClientID TODO
type ClientID = string

// ClientChannel TODO
type ClientChannel = chan Client

// Msg TODO
type Msg = []byte

// Room TODO
type Room struct {
	ID string `json:"id"`

	Title string `json:"title"`

	Clients []ClientID `json:"clients"`

	clients Clients

	// Inbound messages from clients
	broadcast chan Msg

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
		Clients: []ClientID{},

		clients:    Clients{},
		broadcast:  make(chan Msg),
		register:   make(ClientChannel),
		unregister: make(ClientChannel),
	}
}

// Handler TODO
type Handler interface {
	OnRoomChange(room *Room)
	OnRoomJoin(room *Room, client Client)
	OnRoomLeave(room *Room, client Client)
}

// On TODO
func (room *Room) On(handler Handler) {
	for {
		select {
		case client := <-room.register:
			if room.Has(client) {
				return
			}
			room.addClient(client)
			log.Printf("client %v join room %v ", client.ID, room.ID)

			handler.OnRoomJoin(room, client)
			handler.OnRoomChange(room)

		case client := <-room.unregister:
			if !room.Has(client) {
				return
			}

			room.removeClient(client)
			log.Printf("client %v leave room %v ", client.ID, room.ID)

			if !client.Closed {
				handler.OnRoomLeave(room, client)
			}
			handler.OnRoomChange(room)

		case msg := <-room.broadcast:
			for client := range room.clients {
				client.Send(msg)
			}
		}
	}
}

// Broadcast TODO
func (room *Room) Broadcast(msg Msg) {
	room.broadcast <- msg
}

// Join TODO
func (room *Room) Join(client Client) {
	room.register <- client
}

// Leave TODO
func (room *Room) Leave(client Client) {
	room.unregister <- client
}

// Has TODO
func (room *Room) Has(target Client) bool {
	return room.clients[target]
}

func (room *Room) addClient(client Client) {
	room.Clients = append(room.Clients, client.ID)
	room.clients[client] = true
	client.RoomID = room.ID
}

func (room *Room) removeClient(client Client) {
	room.Clients = utils.Remove(room.Clients, client.ID)
	delete(room.clients, client)
	client.RoomID = ""
}
