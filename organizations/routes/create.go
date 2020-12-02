package routes

import (
	"opus-cm/organizations/events"
	"opus-cm/organizations/exceptions"
	"opus-cm/organizations/models"
	"opus-cm/organizations/nats"
	"opus-cm/organizations/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

// CreateOrganization creates an organization and saves it to the database.
func CreateOrganization(ctx *fiber.Ctx) error {
	data := models.OrganizationCreate{}
	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.BadRequestError(ctx, exceptions.NewErrorDetail("Bad Request Error", "Could not parse body"))
	}
	if err := validation.Validate(data, ctx); err != nil {
		return exceptions.BadRequestError(ctx, err...);
	}

	ID, _ := uuid.NewV4()
	org := models.Organization{
		ID:      ID.String(),
		Name:    "this is the org name",
		Version: 1,
	}
	publisher := events.Publisher{
		Subject: "test",
		Client:  nats.GetClient(ctx),
		Object:  org,
	}
	publisher.Publish()
	return ctx.Status(201).JSON(org)
}
