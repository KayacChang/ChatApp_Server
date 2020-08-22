package client

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"server/event"
	"server/model"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Handler TODO
type Handler interface {
	OnEvent(evt event.Event, client *Client)
	OnClose(client *Client)
}

// Client TODO
type Client struct {
	ID     string `json:"id"`
	RoomID string `json:"-"`

	ctx    context.Context
	cancel context.CancelFunc
	conn   *websocket.Conn
	send   chan []byte
}

// New TODO
func New(conn *websocket.Conn) *Client {
	ctx, cancel := context.WithCancel(context.Background())
	return &Client{
		ctx:    ctx,
		cancel: cancel,
		conn:   conn,
		send:   make(chan []byte, 256),
	}
}

// On TODO
func (client *Client) On(handler Handler) {
	go client.read(handler)
	go client.write()
}

// Close TODO
func (client *Client) Close() {
	client.cancel()
	close(client.send)
	log.Printf("client id: %v connection close...\n", client.ID)
}

// Send TODO
func (client *Client) Send(msg []byte) {
	client.send <- msg
}

func (client *Client) read(handler Handler) {
	defer client.Close()

	setup(client.conn)

	for {
		select {

		case <-client.ctx.Done():
			handler.OnClose(client)

			return

		default:
			_, bytes, err := client.conn.ReadMessage()

			if err != nil {
				if isUnexpectedCloseError(err) {
					log.Printf("error: %v", err)
				}

				handler.OnClose(client)

				return
			}

			evt := event.Event{}

			if err := json.Unmarshal(bytes, &evt); err != nil {
				log.Printf("error: %v", err)
				break
			}

			handler.OnEvent(evt, client)
		}
	}
}

func (client *Client) write() {
	ticker := time.NewTicker(pingPeriod)

	for {
		select {

		case <-client.ctx.Done():
			ticker.Stop()
			return

		case msg, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				client.conn.WriteMessage(model.Close, []byte{})
				return
			}

			w, err := client.conn.NextWriter(model.Text)
			if err != nil {
				return
			}
			w.Write(msg)
			w.Close()

		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			client.conn.WriteMessage(model.Ping, nil)
		}
	}
}

func setup(client *websocket.Conn) {
	client.SetReadLimit(maxMessageSize)

	client.SetReadDeadline(time.Now().Add(pongWait))

	client.SetPongHandler(func(string) error {
		return client.SetReadDeadline(time.Now().Add(pongWait))
	})
}

func isUnexpectedCloseError(err error) bool {
	return websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
}

func preprocess(msg []byte) []byte {
	return bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
}
