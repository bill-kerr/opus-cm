package exceptions

import "github.com/gofiber/fiber/v2"

// InternalServerError returns a 500 response due to an unknown server error.
func InternalServerError(ctx *fiber.Ctx) error {
	detail := NewErrorDetail("Internal Server Error", "An unkown error occurred.")
	response := NewErrorResponse(detail, 500)
	return ctx.Status(response.StatusCode).JSON(response)
}

// BadRequestError returns a 400 error along with associated details using the current request context.
func BadRequestError(ctx *fiber.Ctx, details ...ErrorDetail) error {
	if len(details) == 1 {
		if details[0].Name == "" {
			details[0].Name = "Bad Request Error"
		}
		response := NewErrorResponse(details[0], 400)
		return ctx.Status(response.StatusCode).JSON(response)
	}
	response := NewMultiErrorResponse("Bad Request Error", 400, details)
	return ctx.Status(response.StatusCode).JSON(response)
}

// InsufficientPermissionsError returns a 403 error indicating that the user does not have sufficient permissions to perform the operation.
func InsufficientPermissionsError(ctx *fiber.Ctx) error {
	detail := NewErrorDetail("Insufficient Permissions Error", "You do not have the requisite permissions to perform this operation.")
	response := NewErrorResponse(detail, 403)
	return ctx.Status(response.StatusCode).JSON(response)
}

// UnauthorizedError returns a 401 error indicating that an Authorization error occurred during the operation.
func UnauthorizedError(ctx *fiber.Ctx, message string) error {
	detail := NewErrorDetail("Unauthorized Error", message)
	response := NewErrorResponse(detail, 401)
	return ctx.Status(response.StatusCode).JSON(response)
}

// NotFoundError returns a 404 error indicating that the requested resource was not found.
func NotFoundError(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = "The requested resource was not found."
	}
	detail := NewErrorDetail("Not Found Error", message)
	response := NewErrorResponse(detail, 404)
	return ctx.Status(response.StatusCode).JSON(response)
}
