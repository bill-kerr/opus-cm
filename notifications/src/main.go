package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	stan "github.com/nats-io/stan.go"
)

func main() {
	app := fiber.New()

	sc, err := stan.Connect("opuscm", "notifications", stan.NatsURL("http://nats-srv:4222"))
	if err != nil {
		log.Fatal(err.Error() + "\n")
	}
	log.Println("Connected to NATS.")

	app.Get("/", func(c *fiber.Ctx) error {
		sc.PublishAsync("test", []byte("Hello World"), nil)
		return c.JSON(fiber.Map{"message": "this is the root route of the notifications service"})
	})

	app.Listen(":3000")
}