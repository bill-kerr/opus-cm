package events

import (
	"opus-cm/organizations/src/nats"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
)

// Listener is the struct for listening for events
type Listener struct {
	Subject        string
	QueueGroupName string
	AckWait        time.Duration
	Client         *nats.Client
	Handler        MessageHandler
}

// MessageHandler is the type for handling messages received from subscriptions.
type MessageHandler interface {
	Parse(msg *stan.Msg)
}

// SubscriptionOptions returns SubscriptionOptions as defined by the struct properties.
func (l *Listener) SubscriptionOptions() []stan.SubscriptionOption {
	return []stan.SubscriptionOption{
		stan.SetManualAckMode(),
		stan.DurableName(l.QueueGroupName),
		stan.AckWait(l.AckWait),
		stan.StartAt(pb.StartPosition_First),
	}
}

// Listen subscribes to the listener's subject
func (l *Listener) Listen() {
	l.Client.SC.QueueSubscribe(l.Subject, l.QueueGroupName, l.onMessage, l.SubscriptionOptions()...)
}

func (l *Listener) onMessage(msg *stan.Msg) {
	l.Handler.Parse(msg)
}
