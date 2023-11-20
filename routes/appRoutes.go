package routes

import (
	"qr_code/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "api is working",
		})
	})

	qrEdnpoint := app.Group("/qr")
	qrEdnpoint.Post("/generate-qr", controller.GenerateQR)
	qrEdnpoint.Get("/generate-qr-logo", controller.GenerateQRWithLogo)
}
