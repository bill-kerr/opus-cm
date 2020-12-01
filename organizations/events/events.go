package events

// Event represents a message sent or received through NATS.
type Event struct {
	Subject string
	Data interface{}
}