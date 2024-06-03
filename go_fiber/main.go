package main

import (
	"go_fiber/route"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	route.AddRoute(app)
	app.Listen(":3000")
}
