package auth

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// RequireAuth is a Fiber middleware for requiring a verified ID token on a request.
func RequireAuth(ctx *fiber.Ctx) error {
	token, err := extractBearerToken(ctx)
	if err != nil {
		return err
	}
	if err = VerifyToken(token); err != nil {
		return err
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
