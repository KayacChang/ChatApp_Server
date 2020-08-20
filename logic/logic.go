package logic

import (
	"encoding/json"
	"log"
	"server/event"
	"server/model"
	"server/room"
)

var rooms map[string]model.Room

func init() {
	room := room.New()

	go room.Start()

	rooms = map[string]model.Room{
		"01": room,
	}
}

// Handle TODO
func Handle(client model.Client) {

	client.On(func(evt event.Event) {
		room := rooms["01"]

		switch evt.Type {

		case event.Room:
			if evt.Action == event.Join {
				room.Add(client)
			}

			if evt.Action == event.Leave {
				room.Delete(client)
			}

		case event.Msg:
			if evt.Action == event.Send && room.Has(client) {

				evt.Action = event.Get

				bytes, err := json.Marshal(evt)
				if err != nil {
					log.Fatalf("error: %v", err)
				}

				room.Broadcast(bytes)
			}
		}
	})
}
