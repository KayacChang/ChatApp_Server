package room

import (
	"log"
	"server/utils"
)

// ClientID TODO
type ClientID = string

// ClientChannel TODO
type ClientChannel = chan ClientID

// Msg TODO
type Msg = []byte

// Room TODO
type Room struct {
	ID string `json:"id"`

	Title string `json:"title"`

	Clients []ClientID `json:"clients"`

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

		broadcast:  make(chan Msg),
		register:   make(ClientChannel),
		unregister: make(ClientChannel),
	}
}

// Handler TODO
type Handler interface {
	OnChange(room *Room)
	OnMsg(room *Room, msg Msg)
}

// On TODO
func (room *Room) On(handler Handler) {
	for {
		select {
		case client := <-room.register:
			room.addClient(client)
			log.Printf("client %v join room %v ", client, room.ID)
			handler.OnChange(room)

		case client := <-room.unregister:
			if !room.Has(client) {
				return
			}

			room.removeClient(client)
			log.Printf("client %v leave room %v ", client, room.ID)
			handler.OnChange(room)

		case msg := <-room.broadcast:
			handler.OnMsg(room, msg)
		}
	}
}

// Broadcast TODO
func (room *Room) Broadcast(msg Msg) {
	room.broadcast <- msg
}

// Join TODO
func (room *Room) Join(client ClientID) {
	room.register <- client
}

// Leave TODO
func (room *Room) Leave(client ClientID) {
	room.unregister <- client
}

// Has TODO
func (room *Room) Has(target ClientID) bool {
	for _, client := range room.Clients {
		if client == target {
			return true
		}
	}
	return false
}

func (room *Room) addClient(client ClientID) {
	room.Clients = append(room.Clients, client)
}

func (room *Room) removeClient(client ClientID) {
	room.Clients = utils.Remove(room.Clients, client)
}
