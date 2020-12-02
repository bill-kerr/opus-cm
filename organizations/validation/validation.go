package validation

import (
	"opus-cm/organizations/exceptions"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validate runs the validation rules defined on the model and returns any errors via the request context.
func Validate(s interface{}, ctx *fiber.Ctx) []exceptions.ErrorDetail {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		details := []exceptions.ErrorDetail{}
		for _, err := range err.(validator.ValidationErrors) {
			detail := exceptions.NewErrorDetail("Validation error", err.ActualTag())
			details = append(details, detail)
		}
		return details
	}
	return nil
}