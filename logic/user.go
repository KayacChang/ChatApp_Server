package logic

import (
	"encoding/json"
	"log"
	"server/event"
	"time"

	"github.com/google/uuid"
)

type userReq struct {
	Name string `json:"username"`
}

type userResp struct {
	Name    string `json:"username"`
	Message string `json:"message"`
}

func (logic *Logic) onUserJoin(user userReq, client Client) {
	client.ID = user.Name
	logic.clients = append(logic.clients, client)

	data, err := json.Marshal(event.Event{
		Type:   event.User,
		Action: event.Join,
		Status: event.OK,
		Data: userResp{
			Name:    client.ID,
			Message: "User Join Success",
		},
	})

	if err != nil {
		log.Printf("error: %v", err)

		return
	}

	client.Send(data)

	broadcastRoomStatus(Clients{client}, logic.rooms)
}

func (logic *Logic) onUserLeave(client Client) {
	logic.OnClose(client)

	data, err := json.Marshal(event.Event{
		Type:   event.User,
		Action: event.Leave,
		Status: event.OK,
		Data: userResp{
			Name:    client.ID,
			Message: "User Leave Success",
		},
	})

	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client.Send(data)
}

type msgReq struct {
	From    string `json:"from"`
	Message string `json:"message"`
}

func (logic *Logic) onUserSendMsg(msg msgReq, client Client) {
	room := findRoomByID(logic.rooms, client.RoomID)
	if room == nil {
		log.Printf("can not find room by id: %v", client.RoomID)

		return
	}

	if !room.Has(client) {
		log.Printf("can not find client %v in room id %v", client.ID, client.RoomID)

		return
	}

	room.Broadcast(Msg{
		ID:      uuid.New().String(),
		Name:    msg.From,
		Message: msg.Message,
		Time:    time.Now(),
	})
}
