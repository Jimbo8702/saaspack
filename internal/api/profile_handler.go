package api

import (
	"Jimbo8702/saaspack/internal/db"
	"Jimbo8702/saaspack/internal/logger"
	"Jimbo8702/saaspack/internal/types"
	"Jimbo8702/saaspack/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	store 		db.ProfileStore
	logger 		logger.Logger
	validator 	validator.Validator
}

func NewProfileHandler(s db.ProfileStore, l logger.Logger, v validator.Validator) *ProfileHandler {
	return &ProfileHandler{
		store: s,
		logger: l,
		validator: v,
	}
}

func (h *ProfileHandler) HandlePostProfile(c *fiber.Ctx) error {
	var params types.CreateProfileParams
	if err := c.BodyParser(&params); err != nil {
		h.logger.Log("error", "creating profile", err)
		return ErrBadRequest()
	}
	profile := types.NewProfileFromParams(params)
	insertedProfile, err := h.store.InsertProfile(c.Context(), profile)
	if err != nil {
		return err
	}
	h.logger.Log("info", "profile inserted", insertedProfile.ID)
	return c.JSON(insertedProfile)
}

func (h *ProfileHandler) HandleGetProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	profile, err := h.store.GetProfileById(c.Context(), id)
	if err != nil {
		h.logger.Log("error", "couldn't get profile", id)
		return ErrResourceNotFound("profile")
	}
	h.logger.Log("info", "profile found", id)
	return c.JSON(profile)
}

func (h *ProfileHandler) HandleListProfile(c *fiber.Ctx) error {
	profiles, err := h.store.ListProfiles(c.Context(), db.Map{})
	if err != nil {
		h.logger.Log("error", "couldn't list categories", nil)
		return ErrResourceNotFound("profiles")
	}
	h.logger.Log("info", "profiles found", nil)
	return c.JSON(profiles)
}

func (h *ProfileHandler) HandleDeleteProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.DeleteProfile(c.Context(), id); err != nil {
		h.logger.Log("error", "couldn't delete profile", id)
		return err
	}
	h.logger.Log("info", "profile deleted", id)
	return c.JSON(DeleteResponse(id))
}