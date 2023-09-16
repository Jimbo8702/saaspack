package api

import (
	"github.com/gofiber/fiber/v2"
)

func AddProductRoutes(app fiber.Router, h *ProductHandler) *fiber.Router {
	app.Post("/products", h.HandlePostProduct)
	app.Delete("/products/:id", h.HandleDeleteProduct)
	app.Get("/products", h.HandleListProducts)
	app.Get("/products/:id", h.HandleGetProduct)
	return &app
}

func AddCategoryRoutes(app fiber.Router, h *CategoryHandler) *fiber.Router {
	app.Post("/category", h.HandlePostCategory)
	app.Delete("/category/:id", h.HandleDeleteCategory)
	app.Get("/category", h.HandleListCategory)
	app.Get("/category/:id", h.HandleGetCategory)
	return &app
}

