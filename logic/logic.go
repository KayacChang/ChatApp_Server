package logic

import (
	"encoding/json"
	"log"
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
	}, func() {
		logic.clients = remove(logic.clients, client)

		room := findRoomByID(logic.rooms, client.RoomID)
		if room != nil {
			room.Leave() <- client
		}
	})
}

func remove(clients Clients, target Client) Clients {
	res := Clients{}

	for _, client := range clients {
		if client != target {
			res = append(res, client)
		}
	}

	return res
}

// InitRoomStatusBroadcast TODO
func (logic *Logic) InitRoomStatusBroadcast() {
	for _, room := range logic.rooms {
		go room.Start(func(room Room) {
			broadcastRoomStatus(logic.clients, logic.rooms)
		})
	}
}

func toRoomStatusData(rooms Rooms) *[]byte {
	data, err := json.Marshal(event.Event{
		Type:    event.Room,
		Action:  event.Update,
		From:    event.Server,
		Message: rooms,
	})
	if err != nil {
		log.Printf("error: %v", err)

		return nil
	}

	return &data
}

func broadcastRoomStatus(clients Clients, rooms Rooms) {
	data := toRoomStatusData(rooms)
	if data == nil {
		return
	}

	broadcast(clients, *data)
}

func broadcast(clients Clients, msg []byte) {
	for _, client := range clients {
		client.Send(msg)
	}
}

func findRoomByID(rooms Rooms, id string) Room {
	for _, room := range rooms {
		if room.ID == id {
			return room
		}
	}
	return nil
}
