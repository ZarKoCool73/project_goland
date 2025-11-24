package controllers

import (
	"talenthouse/go-api/services"

	"github.com/gofiber/fiber/v2"
)

func QRHandler(c *fiber.Ctx) error {
	var A [][]float64

	// Parsear el body
	if err := c.BodyParser(&A); err != nil || len(A) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Verifica que todas las filas tengan la misma longitud
	columns := len(A[0])
	for _, row := range A {
		if len(row) != columns {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "All rows must have the same length",
			})
		}
	}

	// Llamar a la función QRFactorization
	Q, R, err := services.QRFactorization(A)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"matrixQ": Q,
		"matrixR": R,
	})
}
