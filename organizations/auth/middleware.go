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
		return err
	}

	ctx.Locals("user_id", authToken.UID)
	if role, ok := authToken.Claims["role"]; ok {
		ctx.Locals("role", role)
	}

	return ctx.Next()
}

// RequireAdmin is a middleware that requires the SYS_ADMIN role to proceed.
func RequireAdmin(ctx *fiber.Ctx) error {
	role := ctx.Locals("role").(string)
	if role != "SYS_ADMIN" {
		return exceptions.InsufficientPermissionsError(ctx)
	}
	return ctx.Next()
}

func extractBearerToken(ctx *fiber.Ctx) (string, error) {
	header := ctx.Get("Authorization")
	if len(header) > 7 && strings.ToUpper(header[0:7]) == "BEARER " {
		return header[7:], nil
	}
	return "", errors.New("no 'Bearer ' prefix found")
}
