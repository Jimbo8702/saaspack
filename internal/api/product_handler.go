package api

import (
	"Jimbo8702/saaspack/internal/db"
	"Jimbo8702/saaspack/internal/logger"
	"Jimbo8702/saaspack/internal/types"
	"Jimbo8702/saaspack/internal/validator"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	store 		db.ProductStore
	logger 		logger.Logger
	validator 	validator.Validator
}

func NewProductHandler(s db.ProductStore, l logger.Logger, v validator.Validator) *ProductHandler {
	return &ProductHandler{
		store: s,
		logger: l,
		validator: v,
	}
}

func (h *ProductHandler) HandlePostProduct(c *fiber.Ctx) error {
	var params types.CreateProductParams
	if err := c.BodyParser(&params); err != nil {
		h.logger.Log("error", "creating product", err)
		return fmt.Errorf("error creating product dummy")
	}
	if errors := h.validator.Validate(params); errors != nil {
		return c.JSON(errors)
	}
	product := types.NewProductFromParams(params)
	insertedProduct, err := h.store.InsertProduct(c.Context(), product)
	if err != nil {
		return err
	}
	h.logger.Log("info", "product created", product.ID)
	return c.JSON(insertedProduct)
}




