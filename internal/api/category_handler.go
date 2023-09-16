package api

import (
	"Jimbo8702/saaspack/internal/db"
	"Jimbo8702/saaspack/internal/logger"
	"Jimbo8702/saaspack/internal/types"
	"Jimbo8702/saaspack/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	store 		db.CategoryStore
	logger 		logger.Logger
	validator 	validator.Validator
}

type CreateCategoryRequest struct {
	Category string `json:"category"`
}

func NewCategoryHandler(s db.CategoryStore, l logger.Logger, v validator.Validator) *CategoryHandler {
	return &CategoryHandler{
		store: s,
		logger: l,
		validator: v,
	}
}

func (h *CategoryHandler) HandlePostCategory(c *fiber.Ctx) error {
	var params CreateCategoryRequest
	if err := c.BodyParser(&params); err != nil {
		h.logger.Log("error", "creating category", err)
		return ErrBadRequest()
	}
	category := types.NewCategory(params.Category)
	insertedCategory, err := h.store.InsertCategory(c.Context(), category)
	if err != nil {
		return err
	}
	h.logger.Log("info", "category inserted", category.ID)
	return c.JSON(insertedCategory)
}

func (h *CategoryHandler) HandleGetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	category, err := h.store.GetCategoryById(c.Context(), id)
	if err != nil {
		h.logger.Log("error", "couldn't get category", id)
		return ErrResourceNotFound("category")
	}
	h.logger.Log("info", "category found", id)
	return c.JSON(category)
}

func (h *CategoryHandler) HandleListCategory(c *fiber.Ctx) error {
	categories, err := h.store.ListCategories(c.Context(), db.Map{})
	if err != nil {
		h.logger.Log("error", "couldn't list categorys", nil)
		return ErrResourceNotFound("category")
	}
	h.logger.Log("info", "categories found", nil)
	return c.JSON(categories)
}

func (h *CategoryHandler) HandleDeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.DeleteCategory(c.Context(), id); err != nil {
		h.logger.Log("error", "couldn't delete category", id)
		return err
	}
	h.logger.Log("info", "category deleted", id)
	return c.JSON(DeleteResponse(id))
}