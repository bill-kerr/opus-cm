package nats

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

// Client is the wrapper for the NATS Streaming Server implementation.
type Client struct {
	SC stan.Conn
}

func (c *Client) onConnect() {
	fmt.Println("Organizations service connected to NATS.")
}

func (c *Client) onClose() {
	fmt.Println("Organizations service NATS connection closed.")
}

// Connect is the function that users must run to connect to the NATS Streaming Server service
func (c *Client) Connect() {
	sc, err := stan.Connect("opuscm", "organizations", stan.NatsURL("http://nats-srv:4222"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	c.SC = sc
	c.onConnect()
}