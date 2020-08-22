package logic

import (
	"encoding/json"
	"log"
	"server/event"
)

// OnRoomJoin TODO
func (logic *Logic) OnRoomJoin(room Room, client Client) {
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

// OnRoomLeave TODO
func (logic *Logic) OnRoomLeave(room Room, client Client) {
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
