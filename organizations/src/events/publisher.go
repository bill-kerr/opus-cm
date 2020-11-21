package events

import (
	"encoding/json"
	"fmt"
	"opus-cm/organizations/src/nats"
)

// Publisher represents an object that can publish events to the NATS server.
type Publisher struct {
	Subject        string
	Client         *nats.Client
	Serializer EventSerializer
}

// EventSerializer is the interface that must be implemented in order to publish events.
type EventSerializer interface {
	Serialize() ([]byte, error)
}

// Publish publishes the event provided by the EventSerializer to the NATS service.
func (p *Publisher) Publish() {
	data, err := p.Serializer.Serialize()
	if err != nil {
		fmt.Println("Error publishing message.")
		return
	}
	p.Client.SC.Publish(p.Subject, data)
}

// OrganizationCreatedSerializer serializes organization:created events
type OrganizationCreatedSerializer struct {
	OrganizationName string
}

// NOTE: can just implement Serialize on the model to avoid having to create a seperate interface
// Serialize implements the EventSerializer interface
func (p OrganizationCreatedSerializer) Serialize() ([]byte, error) {
	type Organization struct {
		Name string `json:"name"`s
	}
	org := Organization{
		Name: "test_org",
	}
	return json.Marshal(org)
}

var testClient = &nats.Client{}
var test Publisher = Publisher{
	Subject: "test",
	Serializer: OrganizationCreatedSerializer{},
	Client: testClient,
}