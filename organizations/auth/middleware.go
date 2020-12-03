package auth

import (
	"errors"
	"opus-cm/organizations/exceptions"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// RequireAuth is a Fiber middleware for requiring a verified ID token on a request.
func RequireAuth(ctx *fiber.Ctx) error {
	token, err := extractBearerToken(ctx)
	if err != nil {
		return err
	}

	authToken, err := VerifyToken(token)
	if err != nil {
		return exceptions.UnauthorizedError(ctx, "The provided authentication token is not valid.")
	}

	ctx.Locals("user_id", authToken.UID)
	if admin, ok := authToken.Claims["admin"]; ok {
		ctx.Locals("admin", admin)
	}

	return ctx.Next()
}

// RequireAdmin is a middleware that requires that the user be an admin to proceed.
func RequireAdmin(ctx *fiber.Ctx) error {
	admin := ctx.Locals("admin").(bool)
	if !admin {
		return exceptions.InsufficientPermissionsError(ctx)
	}
	return ctx.Next()
}

func extractBearerToken(ctx *fiber.Ctx) (string, error) {
	header := ctx.Get("Authorization")
	if header == "" {
		return "", errors.New("The Authorization header must be set.")
	}
	if len(header) > 7 && strings.ToLower(header[0:7]) == "bearer " {
		return header[7:], nil
	}
	return "", errors.New("The Authorization header must be formatted as 'Bearer <token>' where <token> is a valid auth key.")
}
