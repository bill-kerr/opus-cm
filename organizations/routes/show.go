package routes

import (
	"fmt"
	"opus-cm/organizations/exceptions"
	"opus-cm/organizations/models"

	"github.com/gofiber/fiber/v2"
)

// ShowOrganization returns a single organization matching the URL id parameter.
func ShowOrganization(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	org, err := models.FindOrganization(models.Organization{ID: ID})
	if err != nil {
		message := fmt.Sprintf("An organization with an id of %s does not exist.", ID)
		return exceptions.NotFoundError(ctx, message)
	}
	return ctx.Status(200).JSON(org)
}
