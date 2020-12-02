package exceptions

import "github.com/gofiber/fiber/v2"

// InternalServerError returns a 500 response due to an unknown server error.
func InternalServerError(ctx *fiber.Ctx) error {
	detail := NewErrorDetail("Internal server error", "An unkown error occurred.")
	response := NewErrorResponse(detail, 500)
	return ctx.Status(response.StatusCode).JSON(response)
}

// BadRequestError returns a 400 error along with associated details using the current request context.
func BadRequestError(ctx *fiber.Ctx, details ...ErrorDetail) error {
	if len(details) == 1 {
		if details[0].Name == "" {
			details[0].Name = "Bad request"
		}
		response := NewErrorResponse(details[0], 400)
		return ctx.Status(response.StatusCode).JSON(response)
	}
	response := NewMultiErrorResponse("Bad request", 400, details)
	return ctx.Status(response.StatusCode).JSON(response)
}

