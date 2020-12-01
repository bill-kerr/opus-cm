package events

import (
	"fmt"
	"opus-cm/organizations/nats"
)

// Publisher represents an object that can publish events to the NATS server.
type Publisher struct {
	Subject        string
	Client         *nats.Client
	Object         EventSerializer
}

// EventSerializer is the interface that must be implemented in order to publish events.
type EventSerializer interface {
	Serialize() ([]byte, error)
}

// Publish publishes the event provided by the EventSerializer to the NATS service.
func (p *Publisher) Publish() {
	data, err := p.Object.Serialize()
	if err != nil {
		fmt.Println("Error publishing message.")
		return
	}
	p.Client.SC.Publish(p.Subject, data)
}