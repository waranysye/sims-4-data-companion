package handler

import (
	"sims-daas/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CareerHandler struct {
	Usecase usecase.CareerUsecaseInterface // 👈 Pakai Interface juga
}

func NewCareerHandler(u usecase.CareerUsecaseInterface) *CareerHandler {
	return &CareerHandler{Usecase: u}
}

func (h *CareerHandler) FetchRecommendations(c *fiber.Ctx) error {
	data, err := h.Usecase.GetRecommendations()
	if err != nil {
		return err // 👈 Cukup return error-nya, biar Global Error Handler yang mengurus formatnya!
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func (h *CareerHandler) CreateRecommendation(c *fiber.Ctx) error {
	var req usecase.RecommendationRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.Usecase.CreateRecommendation(req); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Recommendation created successfully",
	})
}

func (h *CareerHandler) UpdateRecommendation(c *fiber.Ctx) error {
	careerID, err := strconv.Atoi(c.Params("career_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid career_id")
	}

	traitID, err := strconv.Atoi(c.Params("trait_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid trait_id")
	}

	var req usecase.UpdateRecommendationRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.Usecase.UpdateRecommendation(careerID, traitID, req); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Recommendation updated successfully",
	})
}

func (h *CareerHandler) DeleteRecommendation(c *fiber.Ctx) error {
	careerID, err := strconv.Atoi(c.Params("career_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid career_id")
	}

	traitID, err := strconv.Atoi(c.Params("trait_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid trait_id")
	}

	if err := h.Usecase.DeleteRecommendation(careerID, traitID); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Recommendation deleted successfully",
	})
}
