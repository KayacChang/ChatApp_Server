package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"server/event"
)

func (logic *Logic) onRoomJoin(evt event.Event, client Client) {
	room := findRoomByID(logic.rooms, fmt.Sprintf("%v", evt.Message))
	if room == nil {
		log.Printf("can not find room by id: %v", evt.Message)

		return
	}

	room.Join(client.ID)
	client.RoomID = room.ID

	data, err := json.Marshal(event.Event{
		Type:    event.Room,
		Action:  event.Join,
		From:    event.Server,
		Message: "Room Join Success",
	})

	if err != nil {
		log.Printf("error: %v", err)

		return
	}
	client.Send(data)
}

func (logic *Logic) onRoomLeave(evt event.Event, client Client) {
	room := findRoomByID(logic.rooms, client.RoomID)
	if room == nil {
		log.Printf("can not find room by id: %v", client.RoomID)

		return
	}

	room.Leave(client.ID)
	client.RoomID = ""

	data, err := json.Marshal(event.Event{
		Type:    event.Room,
		Action:  event.Leave,
		From:    event.Server,
		Message: "Room Leave Success",
	})

	if err != nil {
		log.Printf("error: %v", err)

		return
	}
	client.Send(data)
}
