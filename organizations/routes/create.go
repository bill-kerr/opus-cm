package routes

import (
	"opus-cm/organizations/events"
	"opus-cm/organizations/exceptions"
	"opus-cm/organizations/models"
	"opus-cm/organizations/nats"
	"opus-cm/organizations/validation"

	"github.com/gofiber/fiber/v2"
)

// CreateOrganization creates an organization and saves it to the database.
func CreateOrganization(ctx *fiber.Ctx) error {
	data := models.OrganizationCreate{}
	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.BadRequestError(ctx, exceptions.NewErrorDetail("Bad Request Error", "Could not parse body"))
	}
	if err := validation.Validate(data, ctx); err != nil {
		return exceptions.BadRequestError(ctx, err...)
	}
	if _, err := models.FindOrganization(models.Organization{Name: data.Name}); err == nil {
		return exceptions.BadRequestError(ctx, exceptions.NewErrorDetail("Bad Request Error", "An organization with that name already exists."))
	}

	org := models.NewOrganization(data.Name)
	if err := org.Save(); err != nil {
		return exceptions.InternalServerError(ctx)
	}

	publisher := events.Publisher{
		Subject: events.SubjectOrganizationCreated,
		Client:  nats.GetClient(ctx),
		Payload: org,
	}
	publisher.Publish()

	return ctx.Status(201).JSON(org)
}
