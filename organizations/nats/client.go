package nats

import (
	"fmt"
	"log"
	"opus-cm/organizations/config"

	"github.com/gofiber/fiber/v2"
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
	conf := config.GetConfig()
	sc, err := stan.Connect(conf.NATSClusterID, conf.NATSClientID, stan.NatsURL(conf.NATSURL))
	if err != nil {
		log.Fatalln(err.Error())
	}
	c.SC = sc
	c.onConnect()
}

// SetClient sets the passed Client on the request context.
func SetClient(client *Client, ctx *fiber.Ctx) {
	ctx.Locals("nats_client", client)
}

// GetClient retrieves the currently set Client from the request context.
func GetClient(ctx *fiber.Ctx) *Client {
	return ctx.Locals("nats_client").(*Client)
}

// ClientProvider is a middleware that sets the NATS client on the current request context.
func ClientProvider(client *Client) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		SetClient(client, ctx)
		return ctx.Next()
	}
}