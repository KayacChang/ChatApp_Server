package logic

import (
	"encoding/json"
	"log"
	"server/event"
)

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

func findClientByID(clients Clients, id string) Client {
	for _, client := range clients {
		if client.ID == id {
			return client
		}
	}
	return nil
}

func (clients Clients) remove(target Client) Clients {
	res := Clients{}

	for _, client := range clients {
		if client != target {
			res = append(res, client)
		}
	}
	return res
}

func (logic *Logic) sendByClientID(id string, msg []byte) {
	client := findClientByID(logic.clients, id)

	if client != nil {
		client.Send(msg)
	}
}
