package event

// Type TODO
type Type string

const (
	// Room TODO
	Room Type = "ROOM"

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

	// Get TODO
	Get = "GET"
)

// Event TODO
type Event struct {
	Type    Type   `json:"type"`
	Action  Action `json:"action"`
	Message string `json:"message"`
}
