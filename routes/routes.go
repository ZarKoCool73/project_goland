package routes

import (
	"talenthouse/go-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/factorizar", controllers.QRHandler)
}
