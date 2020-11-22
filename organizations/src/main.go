package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/stan.go"

	"opus-cm/organizations/src/events"
	"opus-cm/organizations/src/nats"
	"opus-cm/organizations/src/routes"
)

// TODO: need to implement disconnet on SIGTERM/SIGINT to avoid clientID duplicate registrations

func main() {
	app := fiber.New()
	sc := &nats.Client{}
	sc.Connect()

	testHandler := &testHandler{}
	testListener := &events.Listener{
		QueueGroupName: "organizations-service",
		AckWait: time.Duration(5) * time.Second,
		Client: sc,
		Subject: "test",
		Handler: testHandler,
	}
	testListener.Listen()

	app.Use(nats.ClientProvider(sc))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "this is the root route of the organizations service"})
	})

	app.Post("/new", routes.CreateOrganization)

	app.Listen(":3000")
}

type testHandler struct {}

// Parse implements the MessageHandler interface
func (t *testHandler) Parse(msg *stan.Msg) {
	message := string(msg.Data)
	err := msg.Ack()
	if err != nil {
		fmt.Println("Failed to ack message.")
		return
	}
	fmt.Printf("Message received in the organizations service: %s\n", message)
}