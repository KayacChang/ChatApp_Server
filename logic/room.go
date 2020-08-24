package logic

import (
	"encoding/json"
	"log"
	"server/event"
)

type roomReq struct {
	RoomID string `json:"room_id"`
}

type roomResp struct {
	RoomID  string `json:"room_id"`
	Message string `json:"message"`
}

// OnRoomJoin TODO
func (logic *Logic) OnRoomJoin(room Room, client Client) {
	data, err := json.Marshal(event.Event{
		Type:   event.Room,
		Action: event.Join,
		Status: event.OK,
		Data: roomResp{
			RoomID:  room.ID,
			Message: "Room Join Success",
		},
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
		Type:   event.Room,
		Action: event.Leave,
		Status: event.OK,
		Data: roomResp{
			RoomID:  room.ID,
			Message: "Room Leave Success",
		},
	})

	if err != nil {
		log.Printf("error: %v", err)

		return
	}
	client.Send(data)
}
