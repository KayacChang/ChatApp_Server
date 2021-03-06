package event

// Type TODO
type Type string

const (
	// User TODO
	User Type = "USER"

	// Room TODO
	Room = "ROOM"

	// Msg TODO
	Msg = "MSG"
)

// Action TODO
type Action string

const (
	// Join TODO
	Join Action = "JOIN"

	// Leave TODO
	Leave = "LEAVE"

	// Send TODO
	Send = "SEND"

	// Receive TODO
	Receive = "RECEIVE"

	// Update TODO
	Update = "UPDATE"
)

// Status TODO
type Status string

const (
	// OK TODO
	OK Status = "OK"
	// Error TODO
	Error = "ERROR"
)

// Server TODO
const Server = "SERVER"

// Event TODO
type Event struct {
	Type   Type        `json:"type"`
	Action Action      `json:"action"`
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}
