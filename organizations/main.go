package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nats-io/stan.go"
	"gorm.io/gorm"

	"opus-cm/organizations/auth"
	"opus-cm/organizations/config"
	"opus-cm/organizations/database"
	"opus-cm/organizations/events"
	"opus-cm/organizations/models"
	"opus-cm/organizations/nats"
	"opus-cm/organizations/routes"
)

// TODO: need to implement disconnet on SIGTERM/SIGINT to avoid clientID duplicate registrations

func main() {
	c := config.Init()

	db := database.Init(c.PGConnString)
	migrate(db)

	sc := &nats.Client{}
	sc.Connect()

	auth.Init()

	app := fiber.New()
	app.Use(recover.New())
	testHandler := &testHandler{}
	testListener := &events.Listener{
		QueueGroupName: c.QueueGroupName,
		AckWait:        c.AckWait,
		Client:         sc,
		Subject:        "user:created",
		Handler:        testHandler,
	}
	testListener.Listen()

	app.Use(auth.RequireAuth)
	app.Use(auth.RequireAdmin)
	app.Use(nats.ClientProvider(sc))

	app.Get("/", routes.ListOrganizations)
	app.Get("/:id", routes.ShowOrganization)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"test": "this works"})
	})

	app.Post("/", routes.CreateOrganization)

	app.Listen(":3000")
}

type testHandler struct{}

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

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Organization{})
}
