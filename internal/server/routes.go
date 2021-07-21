package server

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, world!")
	})
}
