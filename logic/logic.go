package logic

import (
	"server/client"
	"server/event"
	"server/room"
)

// Client TODO
type Client = *client.Client

// Clients TODO
type Clients = []Client

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
func New(rooms Rooms) *Logic {
	return &Logic{
		rooms:   rooms,
		clients: Clients{},
	}
}

// Handle TODO
func (logic *Logic) Handle(client Client) {

	client.On(func(evt event.Event) {
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
	})
}

func findRoomByID(rooms Rooms, id string) Room {
	for _, room := range rooms {
		if room.ID == id {
			return room
		}
	}
	return nil
}
