package api

import (
	"github.com/gofiber/fiber/v2"
)

func AddProductRoutes(app fiber.Router, h *ProductHandler) *fiber.Router {
	app.Post("/products", h.HandlePostProduct)
	app.Get("/products", h.HandleListProducts)
	app.Get("/products/:id", h.HandleGetProduct)
	//TODO: put
	app.Delete("/products/:id", h.HandleDeleteProduct)
	return &app
}

func AddCategoryRoutes(app fiber.Router, h *CategoryHandler) *fiber.Router {
	app.Post("/category", h.HandlePostCategory)
	app.Get("/category", h.HandleListCategory)
	app.Get("/category/:id", h.HandleGetCategory)
	//TODO: put
	app.Delete("/category/:id", h.HandleDeleteCategory)
	return &app
}

func AddBookingRoutes(app fiber.Router, h *BookingHandler) *fiber.Router {
	app.Post("/booking", h.HandlePostBooking)
	app.Get("/booking/:id", h.HandleGetBooking)
	app.Get("/booking", h.HandleListBooking)
	app.Delete("/booking/:id", h.HandleDeleteBooking)
	return &app
}

func AddProfileRoutes(app fiber.Router, h *ProfileHandler) *fiber.Router {
	app.Post("/profile", h.HandlePostProfile)
	app.Get("/profile", h.HandleListProfile)
	app.Get("/profile/:id", h.HandleGetProfile)
	app.Delete("/profile/:id", h.HandleDeleteProfile)
	return &app
}

