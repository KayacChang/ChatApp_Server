package model

import (
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
