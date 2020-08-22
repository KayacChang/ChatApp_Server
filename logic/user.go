package logic

import (
	"encoding/json"
	"log"
	"server/event"
)

func (logic *Logic) onUserJoin(evt event.Event, client Client) {
	client.ID = evt.From
	logic.clients = append(logic.clients, client)

	data, err := json.Marshal(event.Event{
		Type:    event.User,
		Action:  event.Join,
		From:    event.Server,
		Message: "User Join Success",
	})

	if err != nil {
		log.Printf("error: %v", err)

		return
	}

	client.Send(data)

	broadcastRoomStatus(Clients{client}, logic.rooms)
}

func (logic *Logic) onUserSendMsg(evt event.Event, client Client) {
	room := findRoomByID(logic.rooms, client.RoomID)
	if room == nil {
		log.Printf("can not find room by id: %v", client.RoomID)

		return
	}

	if !room.Has(client.ID) {
		log.Printf("can not find client %v in room id %v", client.ID, client.RoomID)

		return
	}

	bytes, err := json.Marshal(event.Event{
		Type:    event.Msg,
		Action:  event.Receive,
		From:    evt.From,
		Message: evt.Message,
	})

	if err != nil {
		log.Printf("error: %v", err)

		return
	}

	room.Broadcast(bytes)
}
