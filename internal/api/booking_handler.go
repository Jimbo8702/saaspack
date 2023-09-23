package api

import (
	"Jimbo8702/saaspack/internal/db"
	"Jimbo8702/saaspack/internal/logger"
	"Jimbo8702/saaspack/internal/types"
	"Jimbo8702/saaspack/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	store 		db.BookingStore
	logger 		logger.Logger
	validator 	validator.Validator
}

func NewBookingHandler(s db.BookingStore, l logger.Logger, v validator.Validator) *BookingHandler {
	return &BookingHandler{
		store: s,
		logger: l,
		validator: v,
	}
}

func (h *BookingHandler) HandlePostBooking(c *fiber.Ctx) error {
	var params types.CreateBookingParams
	if err := c.BodyParser(&params); err != nil {
		h.logger.Log("error", "creating booking", err)
		return ErrBadRequest()
	}
	booking := types.NewBookingFromParams(params)
	insertedBooking, err := h.store.InsertBooking(c.Context(), booking)
	if err != nil {
		return err
	}
	h.logger.Log("info", "booking inserted", insertedBooking.ID)
	return c.JSON(insertedBooking)
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.GetBookingById(c.Context(), id)
	if err != nil {
		h.logger.Log("error", "couldn't get booking", id)
		return ErrResourceNotFound("booking")
	}
	h.logger.Log("info", "booking found", id)
	return c.JSON(booking)
}

func (h *BookingHandler) HandleListBooking(c *fiber.Ctx) error {
	bookings, err := h.store.ListBookings(c.Context(), db.Map{})
	if err != nil {
		h.logger.Log("error", "couldn't list bookings", nil)
		return ErrResourceNotFound("bookings")
	}
	h.logger.Log("info", "bookings found", nil)
	return c.JSON(bookings)
}

func (h *BookingHandler) HandleDeleteBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.DeleteBooking(c.Context(), id); err != nil {
		h.logger.Log("error", "couldn't delete booking", id)
		return err
	}
	h.logger.Log("info", "category deleted", id)
	return c.JSON(DeleteResponse(id))
}