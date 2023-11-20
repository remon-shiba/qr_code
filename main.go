package main

import (
	"qr_code/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.AppRoutes(app)

	app.Listen(":3333")
}
