package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"talenthouse/go-api/routes"
)

func main() {
	app := fiber.New()

	// Habilitar CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Permite llamadas desde Angular
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
