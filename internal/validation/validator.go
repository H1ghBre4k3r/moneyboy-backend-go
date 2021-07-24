package validation

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type PayloadValidator struct {
	c *fiber.Ctx
}

// Create a new PayloadValidator
func New(c *fiber.Ctx) *PayloadValidator {
	return &PayloadValidator{
		c,
	}
}

// Validate a provided value
func (v *PayloadValidator) Validate(payload interface{}) error {
	if err := v.c.BodyParser(payload); err != nil {
		return err
	}
	return validator.New().Struct(payload)
}
