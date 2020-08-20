package model

import (
	"github.com/gorilla/websocket"
)

const (
	Text  = websocket.TextMessage
	Close = websocket.CloseMessage
)

type Client interface {
	Start(room Room)
	Close()
	Send(msg []byte)
}

type Clients = map[Client]bool

type Room interface {
	Start()
	Broadcast(msg []byte)
	Add(client Client)
	Delete(client Client)
}
