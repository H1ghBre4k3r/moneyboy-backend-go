package server

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, world!")
	})
	auth.New(app)
}
