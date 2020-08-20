package main

import (
	"fmt"
	"log"
	"net/http"
	"server/client"
	"server/logic"

	"github.com/gorilla/websocket"
)

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, "Login")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)

			return
		}

		logic.Handle(client.New(conn))
		log.Println("client connected...")
	})

	log.Printf("server listen on port %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
