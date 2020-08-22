package logic

import (
	"server/client"
	"server/event"
	"server/room"
)

// Client TODO
type Client = *client.Client

// Clients TODO
type Clients []Client

// Room TODO
type Room = *room.Room

// Rooms TODO
type Rooms = []Room

// Logic TODO
type Logic struct {
	rooms   Rooms
	clients Clients
}

// New TODO
func New() *Logic {
	it := &Logic{
		rooms: []Room{
			room.New("01", "Test"),
		},
		clients: []Client{},
	}

	for _, room := range it.rooms {
		go room.On(it)
	}

	return it
}

// Handle TODO
func (logic *Logic) Handle(client Client) {
	client.On(logic)
}

// OnChange TODO
func (logic *Logic) OnChange(room Room) {
	broadcastRoomStatus(logic.clients, logic.rooms)
}

// OnMsg TODO
func (logic *Logic) OnMsg(room Room, msg []byte) {
	for _, id := range room.Clients {
		go logic.sendByClientID(id, msg)
	}
}

// OnEvent TODO
func (logic *Logic) OnEvent(evt event.Event, client Client) {
	switch evt.Type {

	case event.User:
		if evt.Action == event.Join {
			logic.onUserJoin(evt, client)
		}

	case event.Room:
		if evt.Action == event.Join {
			logic.onRoomJoin(evt, client)
		}

		if evt.Action == event.Leave {
			logic.onRoomLeave(evt, client)
		}

	case event.Msg:
		if evt.Action == event.Send {
			logic.onUserSendMsg(evt, client)
		}
	}
}

// OnClose TODO
func (logic *Logic) OnClose(client Client) {
	logic.clients = logic.clients.remove(client)

	room := findRoomByID(logic.rooms, client.RoomID)
	if room != nil {
		room.Leave(client.ID)
	}
}
