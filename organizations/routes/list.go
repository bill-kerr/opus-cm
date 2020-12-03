package routes

import (
	"opus-cm/organizations/exceptions"
	"opus-cm/organizations/models"

	"github.com/gofiber/fiber/v2"
)

// ListOrganizations returns all of the Organizations in the database.
func ListOrganizations(ctx *fiber.Ctx) error {
	orgs, err := models.FindAllOrganizations()
	if err != nil {
		return exceptions.InternalServerError(ctx)
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data":   orgs,
		"object": "list",
	})
}
