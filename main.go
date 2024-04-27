package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(ctx *fiber.Ctx) error {

		return ctx.SendStatus(http.StatusOK)
	})

	app.Listen(":3000")
}
