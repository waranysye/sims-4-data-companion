package handler

import (
	"sims-daas/usecase"

	"github.com/gofiber/fiber/v2"
)

type CareerHandler struct {
	Usecase *usecase.CareerUsecase
}

func NewCareerHandler(u *usecase.CareerUsecase) *CareerHandler {
	return &CareerHandler{Usecase: u}
}

func (h *CareerHandler) FetchRecommendations(c *fiber.Ctx) error {
	data, err := h.Usecase.GetRecommendations()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}
