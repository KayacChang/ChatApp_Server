package main

import (
	"log"
	"net/http"
	"server/client"
	"server/logic"
	"server/room"

	"github.com/gorilla/websocket"
)

// Room TODO
type Room = logic.Room

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	handler := logic.New([]Room{
		room.New("01", "Test"),
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)

			return
		}

		handler.Handle(client.New(conn))
		log.Printf("new connection from %v\n", getIP(r))
	})

	log.Printf("server listen on port %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
