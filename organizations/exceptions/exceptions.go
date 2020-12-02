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

// InsufficientPermissionsError returns a 403 error indicating that the user does not have sufficient permissions to perform the operation.
func InsufficientPermissionsError(ctx *fiber.Ctx) error {
	detail := NewErrorDetail("Insufficient permissions", "You do not have the requisite permissions to perform this operation.")
	response := NewErrorResponse(detail, 403)
	return ctx.Status(response.StatusCode).JSON(response)
}
