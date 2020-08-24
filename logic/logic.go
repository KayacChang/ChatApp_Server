package logic

import (
	"log"
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

// Msg TODO
type Msg = room.Msg

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

// OnRoomChange TODO
func (logic *Logic) OnRoomChange(room Room) {
	broadcastRoomStatus(logic.clients, logic.rooms)
}

// OnEvent TODO
func (logic *Logic) OnEvent(evt event.Event, client Client) {
	switch evt.Type {

	case event.User:

		if evt.Action == event.Join {
			data, ok := evt.Data.(map[string]interface{})
			if !ok {
				log.Printf("can not handle this data: %v\n", evt.Data)

				return
			}

			user := userReq{
				Name: data["username"].(string),
			}
			logic.onUserJoin(user, client)
		}

		if evt.Action == event.Leave {
			logic.onUserLeave(client)
		}

	case event.Room:

		if evt.Action == event.Join {
			data, ok := evt.Data.(map[string]interface{})
			if !ok {
				log.Printf("can not handle this data: %v\n", evt.Data)

				return
			}

			req := roomReq{
				RoomID: data["room_id"].(string),
			}

			room := findRoomByID(logic.rooms, req.RoomID)
			if room == nil {
				log.Printf("can not find room by id: %v", req.RoomID)

				return
			}

			room.Join(client)
		}

		if evt.Action == event.Leave {
			room := findRoomByID(logic.rooms, client.RoomID)
			if room == nil {
				log.Printf("can not find room by id: %v", client.RoomID)

				return
			}

			room.Leave(client)
		}

	case event.Msg:
		data, ok := evt.Data.(map[string]interface{})
		if !ok {
			log.Printf("can not handle this data: %v\n", evt.Data)

			return
		}

		msg := msgReq{
			From:    data["from"].(string),
			Message: data["message"].(string),
		}

		if evt.Action == event.Send {
			logic.onUserSendMsg(msg, client)
		}
	}
}

// OnClose TODO
func (logic *Logic) OnClose(client Client) {
	logic.clients = logic.clients.remove(client)

	room := findRoomByID(logic.rooms, client.RoomID)
	if room != nil {
		room.Leave(client)
	}
}
