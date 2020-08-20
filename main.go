package main

import (
	"log"
	"net/http"
	"server/client"
	"server/room"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	room := room.New()

	go room.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err)

			return
		}

		room.Add(client.New(conn))
	})

	log.Printf("server listen on port %d", 8080)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
