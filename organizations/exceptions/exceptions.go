package exceptions

import "github.com/gofiber/fiber/v2"

// BadRequestError returns a 400 error along with associated details using the current request context.
func BadRequestError(ctx *fiber.Ctx, details ...ErrorDetail) error {
	if len(details) == 1 {
		if details[0].Name == "" {
			details[0].Name = "Bad Request"
		}
		response := NewErrorResponse(details[0], 400)
		return ctx.Status(response.StatusCode).JSON(response)
	}
	response := NewMultiErrorResponse("Bad Request", 400, details)
	return ctx.Status(response.StatusCode).JSON(response)
}