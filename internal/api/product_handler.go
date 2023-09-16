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

//
// POST: create a product
//
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

//
// GET: get a product with a given id
//
func (h *ProductHandler) HandleGetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.store.GetProductById(c.Context(), id)
	if err != nil {
		return ErrResourceNotFound("product")
	}
	return c.JSON(product)
}

//
// LIST (GET MANY): list all products or list products with given filter
//
func (h *ProductHandler) HandleListProducts(c *fiber.Ctx) error {
	//
	//TODO: add the ability to query products by fields
	//
	products, err := h.store.ListProducts(c.Context(), db.Map{})
	if err != nil {
		return ErrResourceNotFound("products")
	}
	return c.JSON(products)
}

func (h *ProductHandler) HandleDeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.DeleteProduct(c.Context(), id); err != nil {
		return err
	}
	return c.JSON(DeleteResponse(id))
}



