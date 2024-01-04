package main

import (
	"qr_code/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	// Initialize default config
	app.Use(requestid.New())

	// Initialize default config
	app.Use(logger.New())
	// Or extend your config for customization
	app.Use(requestid.New(requestid.Config{
		Header: "X-Custom-Header",
		Generator: func() string {
			return "static-id"
		},
	}))

	routes.AppRoutes(app)

	app.Listen(":3333")
}
