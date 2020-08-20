package model

import (
	"server/event"

	"github.com/gorilla/websocket"
)

const (
	// Ping TODO
	Ping = websocket.PingMessage
	// Text TODO
	Text = websocket.TextMessage
	// Close TODO
	Close = websocket.CloseMessage
)

// Listener TODO
type Listener func(evt event.Event)

// Client TODO
type Client interface {
	On(Listener)
	Join(room Room)
	Close()
	Send(msg []byte)
}

// Clients TODO
type Clients = map[Client]bool

// Room TODO
type Room interface {
	Start()
	Broadcast(msg []byte)
	Add(client Client)
	Has(client Client) bool
	Delete(client Client)
}
