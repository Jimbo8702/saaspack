package api

import (
	"github.com/gofiber/fiber/v2"
)

func AddProductRoutes(app fiber.Router, h *ProductHandler) *fiber.Router {
	app.Post("/products", h.HandlePostProduct)
	// app.Get("/products/:id")
	return &app
}